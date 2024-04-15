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

	req := httptest.NewRequest("GET", "/v1/healthcheck", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.healthcheckHandler)
	handler.ServeHTTP(rr, req)

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(body), "status: available") {
		t.Errorf("got %s; expected 'status: available'", string(body))
	}
}
