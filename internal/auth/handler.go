package auth

import (
	"encoding/json"
	"fmt"
	"go/adv-dev/configs"
	"go/adv-dev/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
}
type AuthHandlerDeps struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			res.JsonResponse(w, err.Error(), http.StatusPaymentRequired)
		}
		fmt.Printf("login: email: %s, password: %s\n", payload.Email, payload.Password)
		data := LoginResponse{
			Token: "132",
		}
		res.JsonResponse(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("register")
	}
}
