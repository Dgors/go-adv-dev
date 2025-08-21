package main

import (
	"fmt"
	"go/adv-dev/configs"
	"go/adv-dev/internal/auth"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
