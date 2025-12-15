package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)


func (app *appilication)mount() http.Handler  {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	return  r
}


func (app appilication)run(h http.Handler) error {
	srv := &http.Server{
		Addr: app.config.addr,
		Handler: h,
	}
    
    log.Fatalf("Server started listening on port:%v",app.config.addr)
    return srv.ListenAndServe()
}

type appilication struct{
	config config
}

type config struct{
	addr string
	db   dbConfig
}

type dbConfig struct{
	dsn string
}