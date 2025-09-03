package main

import (
	"fmt"
	"go/adv-dev/configs"
	"go/adv-dev/internal/auth"
	"go/adv-dev/internal/link"
	"go/adv-dev/internal/stat"
	"go/adv-dev/internal/user"
	"go/adv-dev/pkg/db"
	"go/adv-dev/pkg/event"
	"go/adv-dev/pkg/middleware"
	"net/http"
	"time"
)

func main() {
	conf := configs.LoadConfig()
	newDb := db.NewDb(conf)
	router := http.NewServeMux()
	eventBus := event.NewEventBus()

	//Repositories
	linkRepository := link.NewLinkRepository(newDb)
	userRepository := user.NewUserRepository(newDb)
	statRepository := stat.NewStatRepository(newDb)

	//Services
	authService := auth.NewAuthService(userRepository)
	statService := stat.NewStatService(&stat.StatServiceDeps{
		EventBus:       eventBus,
		StatRepository: statRepository,
	})

	//handlers
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
		Config:         conf,
		EventBus:       eventBus,
	})
	stat.NewStatHandler(router, stat.StatHandlerDeps{
		StatRepository: statRepository,
		StatService:    statService,
		Config:         conf,
	})

	// Middlewares
	stack := middleware.Chain(middleware.CORS, middleware.Logging)

	server := http.Server{
		Addr:         ":8081",
		Handler:      stack(router),
		WriteTimeout: 60 * time.Minute,
	}

	go statService.AddClick()

	fmt.Println("server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
