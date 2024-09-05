package user

import (
	"fmt"
	"net/http"

	"github.com/ahenla/go-blog/helpers"
	"github.com/ahenla/go-blog/types"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/signin", h.handleSignIn).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {}

func (h *Handler) handleSignIn(w http.ResponseWriter, r *http.Request) {
	// get the user payload
	var payload types.SignInUserPayload
	if err := helpers.ParseJSON(r, payload); err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err)
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, fmt.Errorf("user with emai %s alredy exists", payload.Email))
	}
	// if does not exists, create a new user
}
