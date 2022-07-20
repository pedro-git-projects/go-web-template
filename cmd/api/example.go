package main

import (
	"fmt"
	"net/http"
	"template/internal/data"
	"template/internal/validator"
	"time"
)

// createExampleHandler will be a POST endpoint
func (app *application) createExampleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string      `json:"title"`
		Year   int32       `json:"year"`
		List   []string    `json:"list"`
		Custom data.Custom `json:"custom"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	example := &data.Example{
		Title:  input.Title,
		Year:   input.Year,
		List:   input.List,
		Custom: input.Custom,
	}

	v := validator.New()

	if data.ValidateExample(*v, *example); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

/*
	interpolatedIDHandler uses the helper readIDParameter to get
	the id from the context of the request and handle it accordingly
*/
func (app *application) interpolatedIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParameter(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	example := data.Example{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "My Cool Example",
		Year:      2020,
		List:      []string{"my", "examples", "are", "cool"},
		Custom:    10,
	}

	err = app.writeJSON(w, http.StatusOK, example, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
