package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/logging"
)

// swagger:route GET /emps employees listEmployees
// Return a list of employees from the database
// responses:
//	200: employeesResponse

// ReturnAllEmps handles GET requests and returns all current employees
func (e *Employees) ReturnAllEmps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	e.l.StandardLogger(logging.Info).Println("Endpoint Hit: returnAllEmps")
	e.l.StandardLogger(logging.Debug).Println(data.Emps)
	json.NewEncoder(w).Encode(data.Emps)

}
