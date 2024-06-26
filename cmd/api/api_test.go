package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"greenlight.doanandreas.net/internal/data"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func Test_HealthCheck(t *testing.T) {
	var app application

	sut := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/healthcheck", nil)

	handler := http.HandlerFunc(app.healthcheckHandler)
	handler.ServeHTTP(sut, r)

	body, err := io.ReadAll(sut.Body)
	if err != nil {
		t.Fatal(err)
	}

	var js healthCheck
	err = json.Unmarshal(body, &js)
	if err != nil {
		t.Fatal(err)
	}

	if js.Status != "available" {
		t.Errorf("got %s; expected 'available'", js.Status)
	}
}

func Test_CreateMovie(t *testing.T) {
	var app application

	sut := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/v1/movies", nil)

	handler := http.HandlerFunc(app.createMovieHandler)
	handler.ServeHTTP(sut, r)

	body, err := io.ReadAll(sut.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), "create a new movie") {
		t.Errorf("got %s; expected 'create a new movie'", string(body))
	}
}

func Test_ShowMovie(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		eStatusCode int
	}{
		{"positive integer", "2", 200},
		{"negative integer", "-3", 404},
		{"non integer", "x13a1g", 404},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var app application

			sut := httptest.NewRecorder()
			r := httptest.NewRequest("GET", "/v1/movies/{id}", nil)

			rCtx := chi.NewRouteContext()
			rCtx.URLParams.Add("id", tt.id)
			r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rCtx))

			handler := http.HandlerFunc(app.showMovieHandler)
			handler.ServeHTTP(sut, r)

			body, err := io.ReadAll(sut.Body)
			if err != nil {
				t.Fatal(err)
			}

			if sut.Result().StatusCode != tt.eStatusCode {
				t.Errorf("got '%d'; expected '%d'", sut.Result().StatusCode, tt.eStatusCode)
			}

			if sut.Result().StatusCode == http.StatusNotFound {
				return
			}

			var js map[string]data.Movie
			err = json.Unmarshal(body, &js)
			if err != nil {
				t.Fatal(err)
			}

			expectedId, err := strconv.ParseInt(tt.id, 10, 64)
			if err != nil && js["movie"].ID != expectedId {
				t.Errorf("got '%d'; expected '%d'", js["movie"].ID, expectedId)
			}
		})
	}
}
