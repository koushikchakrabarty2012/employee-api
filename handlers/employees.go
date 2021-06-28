package handlers

import (
	"net/http"
	"strconv"

	"cloud.google.com/go/logging"
	"github.com/gorilla/mux"
)

// Employees handler for getting and updating employees
type Employees struct {
	l *logging.Logger
}

// NewEmployees returns a new employees handler with the given logger
func NewEmployees(l *logging.Logger) *Employees {
	return &Employees{l}
}

// getEmployeeID returns the employee ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getEmployeeID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
