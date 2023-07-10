package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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

	go func() {
		fmt.Println("Starting server...")
		if err := h.Server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("Shutting down gracefully, press Ctrl+C again to force")
	return nil
}
