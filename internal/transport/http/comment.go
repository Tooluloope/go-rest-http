package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tooluloope/go-rest-http/internal/comment"
)

type CommentService interface {
	PostComment(context.Context, comment.Comment) (comment.Comment, error)
	GetComment(ctx context.Context, ID string) (comment.Comment, error)
	UpdateComment(ctx context.Context, comment comment.Comment) (comment.Comment, error)
	DeleteComment(ctx context.Context, ID string) error
}

type Response struct {
	Message string `json:"message"`
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return
	}

	comment, err := h.Service.PostComment(r.Context(), comment)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		log.Println(err)
		panic(err)
	}

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	cmt, err := h.Service.GetComment(r.Context(), id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		return
	}

	comment.ID = id

	comment, err := h.Service.UpdateComment(r.Context(), comment)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		log.Println(err)
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := h.Service.DeleteComment(r.Context(), id)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(
		Response{
			Message: "Comment deleted successfully",
		},
	); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
