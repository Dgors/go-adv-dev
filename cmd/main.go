package main

import (
	"fmt"
	"go/adv-dev/configs"
	"go/adv-dev/internal/auth"
	"go/adv-dev/internal/link"
	"go/adv-dev/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	newDb := db.NewDb(conf)
	router := http.NewServeMux()

	//Repositories
	linkRepository := link.NewLinkRepository(newDb)

	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{Config: conf})
	link.NewLinkHandler(router, link.LinkHandlerDeps{LinkRepository: linkRepository})

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
