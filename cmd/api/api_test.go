package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_HealthCheck(t *testing.T) {
	var app application

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/v1/healthcheck", nil)

	handler := http.HandlerFunc(app.healthcheckHandler)
	handler.ServeHTTP(w, r)

	sut, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(sut), "status: available") {
		t.Errorf("got %s; expected 'status: available'", string(sut))
	}
}

func Test_CreateMovie(t *testing.T) {
	var app application

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/v1/movies", nil)

	handler := http.HandlerFunc(app.createMovieHandler)
	handler.ServeHTTP(w, r)

	sut, err := io.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(sut), "create a new movie") {
		t.Errorf("got %s; expected 'create a new movie'", string(sut))
	}
}
