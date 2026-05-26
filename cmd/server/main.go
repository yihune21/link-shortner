package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/yihune21/link-shortner/internal/config"
	"github.com/yihune21/link-shortner/internal/handler"
	"github.com/yihune21/link-shortner/internal/service"
	"github.com/yihune21/link-shortner/internal/utils"
)


func main()  {
	err := config.LoadEnv()
    if err != nil {
	   log.Fatalln(err)
	}

	dbUrl := config.GetEnv("GOOSE_DBSTRING")
	if dbUrl == "" {
		log.Fatalln("Database string is empty.") 
	}

	q, err := utils.ConnectDb(dbUrl)

	userService := service.NewUserService(q)
	linkUservice := service.NewLinkService(q)
	authService := service.NewAuthService(q)

	userHandler := handler.NewUserHandler(userService)
	linkHandler := handler.NewLinkHandler(linkUservice)
    authHandler := handler.NewAuthHandler(authService)

    r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            true,
		MaxAge:           300,
	}))

	r.Route("/v1", func(r chi.Router) {
		r.Post("/refresh-token", authHandler.RefreshToken)
		r.Get("/{shortId}", linkHandler.GetLinksByShortLink)
		
		r.Group(func(r chi.Router) {
			r.Use(authHandler.MiddlewareAuth)
			r.Post("/link" , linkHandler.CreateLink)
			r.Post("/logout", userHandler.Logout)
		})

		r.Route("/users", func(r chi.Router) {
			r.Post("/",userHandler.CreateUser)
			r.Post("/login", userHandler.Login)
			r.Get("/", userHandler.ListUser)
			r.Route("/{id}", func(r chi.Router) {
				r.Get("/", userHandler.GetUserById)
				r.Delete("/", userHandler.DeleteUser)
			})
		})
	})
    
	port := config.GetEnv("SERVER_PORT")
	log.Println("Server starting on :" + port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}

