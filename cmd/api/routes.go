package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/car", app.createCarHandler)
	router.HandlerFunc(http.MethodGet, "/car/:id", app.showCarHandler)
	router.HandlerFunc(http.MethodPatch, "/car/:id", app.updateCarHandler)
	router.HandlerFunc(http.MethodDelete, "/car/:id", app.deleteCarHandler)
	router.HandlerFunc(http.MethodGet, "/cars", app.listCarHandler)

	return app.recoverPanic(router)
}
