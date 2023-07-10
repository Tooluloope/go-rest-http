package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CommentService interface{}

type Handler struct {
	Service CommentService
	Router  *mux.Router
	Server  *http.Server
}

func NewHandler(service CommentService) *Handler {
	handler := &Handler{
		Service: service,
	}
	handler.Router = mux.NewRouter()
	handler.mapRoutes()
	handler.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler.Router,
	}
	return handler
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
}

func (h *Handler) Serve() error {
	fmt.Println("Starting server...")
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
