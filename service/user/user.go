package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/signin", h.handleSignIn).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) handleSignIn(w http.ResponseWriter, r *http.Request) {}
