package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// routes creates and returns the handler mux for the application
func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)                 // custom 404 handler
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse) // custom 405 handler

	// Add your routes here
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/example", app.createExampleHandler)
	router.HandlerFunc(http.MethodGet, "/v1/example/:id", app.interpolatedIDHandler)

	return router
}
