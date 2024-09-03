package user

import (
	"net/http"

	"github.com/ahenla/go-blog/helpers"
	"github.com/ahenla/go-blog/types"
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

func (h *Handler) handleSignIn(w http.ResponseWriter, r *http.Request) {
	// get the user payload
	var payload types.SignInUserPayload
	helpers.ParseJSON(r, payload)
	// check if the user exists
	// if does not exists, create a new user
}
