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
		return
	}

	// check if the user exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, fmt.Errorf("user with email %s alredy exists", payload.Email))
		return
	}

	// if does not exists, create a new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  payload.Password,
	})
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, nil)

}
