package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/logging"
)

// swagger:route GET /emps/{id} employees listSingleEmployee
// Return single employee from the database
// responses:
//	200: employeeResponse
//	404: errorResponse
// ReturnSingleEmp handles GET requests
func (e *Employees) ReturnSingleEmp(w http.ResponseWriter, r *http.Request) {
	e.l.StandardLogger(logging.Info).Println("Endpoint Hit: returnSingleEmp")
	w.Header().Set("Content-Type", "application/json")
	id := getEmployeeID(r)
	e.l.StandardLogger(logging.Debug).Println("[DEBUG] get record id", id)
	emp, err := data.GetEmployeeByID(id)
	switch err {
	case nil:

	case data.ErrEmployeeNotFound:
		e.l.StandardLogger(logging.Error).Println("[ERROR] fetching employee", err)

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&data.GenericError{Message: err.Error()})
		return
	default:
		e.l.StandardLogger(logging.Error).Println("[ERROR] fetching employee", err)

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&data.GenericError{Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(emp)
}
