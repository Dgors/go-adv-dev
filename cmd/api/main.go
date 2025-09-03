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

func App() http.Handler {
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

	go statService.AddClick()
	// Middlewares
	stack := middleware.Chain(middleware.CORS, middleware.Logging)

	return stack(router)
}

func main() {
	app := App()
	server := http.Server{
		Addr:         ":8081",
		Handler:      app,
		WriteTimeout: 15 * time.Second,
	}

	fmt.Println("server is listening on port 8081")
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
