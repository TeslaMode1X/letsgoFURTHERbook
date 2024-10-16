package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	r := httprouter.New()

	// httprouter has default responses if some path doesn't exist or some error handling
	// httprouter non-existent url path dealing
	// i'm passing my own json response instead of default plain text
	r.NotFound = http.HandlerFunc(app.notFoundResponse)

	// Likewise, convert the methodNotAllowedResponse() helper to a http.Handler and set
	// it as the custom error handler for 405 Method Not Allowed responses.
	r.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	r.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// Movie creating, showing
	r.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)
	r.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandler)

	return r
}
