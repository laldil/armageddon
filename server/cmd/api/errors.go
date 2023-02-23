package main

import (
	"fmt"
	"net/http"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.PrintError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})

}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(500)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	app.errorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) failedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	app.errorResponse(w, r, http.StatusUnprocessableEntity, errors)
}

func (app *application) invalidCredentialsResponse(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) invalidAuthenticationTokenResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) authenticationRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	app.errorResponse(w, r, http.StatusUnauthorized, message)
}

func (app *application) inactiveAccountResponse(w http.ResponseWriter, r *http.Request) {
	message := "your user account must be activated to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) moderatorRoleRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must have moderator or admin role to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) adminRoleRequiredResponse(w http.ResponseWriter, r *http.Request) {
	message := "you must have admin role to access this resource"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) wrongCarResponse(w http.ResponseWriter, r *http.Request) {
	message := "you can't change the data of someone other than your transport"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) carOccupiedResponse(w http.ResponseWriter, r *http.Request) {
	message := "the car is already occupied"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) carNotUsedResponse(w http.ResponseWriter, r *http.Request) {
	message := "this car is not used"
	app.errorResponse(w, r, http.StatusForbidden, message)
}

func (app *application) fileSizeLimitResponse(w http.ResponseWriter, r *http.Request) {
	message := "the file exceeds the maximum allowed memory size"
	app.errorResponse(w, r, http.StatusNotAcceptable, message)
}
