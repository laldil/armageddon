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

	router.HandlerFunc(http.MethodGet, "/car/:id", app.showCarHandler)
	router.HandlerFunc(http.MethodGet, "/cars", app.listCarHandler)
	router.HandlerFunc(http.MethodPost, "/car", app.requireActivatedUser(app.createCarHandler))
	router.HandlerFunc(http.MethodPatch, "/car/:id", app.requireActivatedUser(app.updateCarHandler))
	router.HandlerFunc(http.MethodDelete, "/car/:id", app.requireActivatedUser(app.deleteCarHandler))

	router.HandlerFunc(http.MethodPut, "/car/:id/rent", app.requireActivatedUser(app.rentCarHandler))
	router.HandlerFunc(http.MethodPut, "/car/:id/return", app.requireActivatedUser(app.returnRentedCarHandler))

	router.HandlerFunc(http.MethodPost, "/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.authenticate(router))
}
