package main

import (
	"fmt"
	"go/adv-dev/configs"
	"go/adv-dev/internal/hello"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	router := http.NewServeMux()
	hello.NewHelloHandler(router)

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
