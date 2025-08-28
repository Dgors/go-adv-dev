package auth

import (
	"fmt"
	"go/adv-dev/configs"
	"go/adv-dev/pkg/req"
	"go/adv-dev/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
}
type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		payload, err := req.HandleBody[LoginRequest](&w, request)
		if err != nil {
			fmt.Println("Error handling body:", err)
			return
		}

		fmt.Printf("login: email: %s, password: %s\n", payload.Email, payload.Password)
		data := LoginResponse{
			Token: "132",
		}
		res.JsonResponse(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, request)
		if err != nil {
			return
		}
		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			res.JsonResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.JsonResponse(w, email, http.StatusCreated)
	}
}
