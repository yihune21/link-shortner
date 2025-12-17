package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/yihune21/link-shortner/internal/database"
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

		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("all good"))
		})

		return  r
}

func (app *appilication)run(h http.Handler) error {
	srv := &http.Server{
		Addr:app.config.addr,
		Handler: h,
	}
    
    log.Fatalf("Server started listening on port:%v",app.config.addr)
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