package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/v1/healthcheck", app.healthcheckHandler)
	mux.Post("/v1/movies", app.createMovieHandler)
	mux.Get("/v1/movies/{id}", app.showMovieHandler)

	return mux
}
