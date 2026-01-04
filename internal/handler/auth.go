package handler

import (
	"encoding/json"
	"mastery-project/internal/model"
	"mastery-project/internal/server"
	"mastery-project/internal/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type AuthHandler struct {
	Handler
	AuthService *service.AuthService
}

func NewAuthHandler(srv *server.Server, auth *service.AuthService) *AuthHandler {
	return &AuthHandler{
		Handler:     NewHandler(srv),
		AuthService: auth,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user model.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validate.Struct(&user); err != nil {
		h.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	response, value, err := h.AuthService.Login(ctx, user)

	if err != nil {
		h.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	cookie := &http.Cookie{
		Name:     "sessionToken",
		Value:    value,
		Path:     "/",
		SameSite: http.SameSiteLaxMode, //mitigate against csrf attacks
		HttpOnly: true,
		Secure:   h.Handler.server.Config.ENV == "production",
	}
	http.SetCookie(w, cookie)

	h.JSON(w, http.StatusOK, response)

}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user model.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := validate.Struct(&user); err != nil {
		h.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	ctx := r.Context()
	response, err := h.AuthService.Register(ctx, user)
	if err != nil {
		h.JSON(w, http.StatusBadRequest, err.Error())
		return
	}
	h.JSON(w, http.StatusOK, response)
}
