package main

import (
	"booking-app/internal/store"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	store store.Storage
}

type config struct {
	addr string
	db dbConfig
}

type dbConfig struct {
	addr string
	maxOpenConns int
	maxIdleConns int
	maxIddleTime string
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Get("/health", app.healthCheckHundler)

	return r
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Minute,
	}

	log.Printf("Server has started at port: %s", app.config.addr)

	return srv.ListenAndServe()
}