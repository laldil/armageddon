package main

import (
	"armageddon/internal/models"
	"armageddon/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createCarHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Brand       string `json:"brand"`
		Description string `json:"description"`
		Color       string `json:"color,omitempty"`
		Year        int32  `json:"year,omitempty"`
		Price       int32  `json:"price"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	car := &models.Car{
		Brand:       input.Brand,
		Description: input.Description,
		Color:       input.Color,
		Year:        input.Year,
		Price:       input.Price,
	}

	v := validator.New()

	if models.ValidateCar(v, car); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Car.Insert(car)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/car/%d", car.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"car": car}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showCarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	car, err := app.models.Car.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"car": car}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateCarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	car, err := app.models.Car.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Brand       *string `json:"brand"`
		Description *string `json:"description"`
		Color       *string `json:"color"`
		Year        *int32  `json:"year"`
		Price       *int32  `json:"price"`
		IsUsed      *bool   `json:"is_used"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Brand != nil {
		car.Brand = *input.Brand
	}

	if input.Description != nil {
		car.Description = *input.Description
	}

	if input.Color != nil {
		car.Color = *input.Color
	}

	if input.Year != nil {
		car.Year = *input.Year
	}

	if input.Price != nil {
		car.Price = *input.Price
	}

	if input.IsUsed != nil {
		car.IsUsed = *input.IsUsed
	}

	v := validator.New()
	if models.ValidateCar(v, car); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Car.Update(car)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"car": car}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Car.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "car successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
