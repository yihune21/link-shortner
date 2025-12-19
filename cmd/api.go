package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/yihune21/link-shortner/internal/database"
	"github.com/yihune21/link-shortner/internal/links"
)


func (app *appilication)mount() http.Handler  {
	    r := chi.NewRouter()

		r.Use(middleware.RequestID)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)

		// Set a timeout value on the request context (ctx), that will signal
		// through ctx.Done() that the request has timed out and further
		// processing should be stopped.
		r.Use(middleware.Timeout(60 * time.Second))
        r.Use(cors.Handler(
			cors.Options{
			AllowedOrigins: []string{"https://*"},
			AllowedMethods: []string{"GET","POST","DELETE","PUT","PATCH"},
			AllowedHeaders: []string{"*"},
			ExposedHeaders: []string{"Link"},
			AllowCredentials: false,
			MaxAge:             300,
		}))
		
		linkHandler := links.NewHandler(links.NewService(app.db))
		r1 := chi.NewRouter()
		
		r1.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("all good"))
		})
		r1.Post("/link",linkHandler.CreateLink)
		r1.Get("/link",linkHandler.ListLinks)
        r1.Get("/link/{id}" , linkHandler.GetLink)
        
		r.Mount("/api",r1)

		return  r
}

func (app *appilication)run(h http.Handler) error {
	srv := &http.Server{
		Addr:":" + app.config.addr,
		Handler: h,
	}
    fmt.Printf("Server listen on port:%v\n",app.config.addr)
    return srv.ListenAndServe()
}

type appilication struct{
	config config
	db *database.Queries
}

type config struct{
	addr string
	db   dbConfig
}

type dbConfig struct{
	dsn string
}