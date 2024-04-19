package main

import (
	"net/http"
)

type healthCheck struct {
	Status     string     `json:"status"`
	SystemInfo systemInfo `json:"system_info"`
}

type systemInfo struct {
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := healthCheck{
		Status: "available",
		SystemInfo: systemInfo{
			Environment: app.config.env,
			Version:     version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
