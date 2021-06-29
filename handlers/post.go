package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/logging"
)

// swagger:route POST /emps employees createEmployee
// Create a new Employee
//
// responses:
//	200: employeeResponse

// Create handles POST requests to add new employees
func (e *Employees) CreateSingleEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee data.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	e.l.StandardLogger(logging.Debug).Println("Requst Payload to the service")
	e.l.Log(logging.Entry{
		// Log anything that can be marshaled to JSON.
		Payload:  employee,
		Severity: logging.Debug,
	})
	employee.ID = data.Count
	data.Count++
	data.Emps = append(data.Emps, employee)
	json.NewEncoder(w).Encode(employee)
	e.l.StandardLogger(logging.Debug).Println("Response Payload to the service")
	e.l.Log(logging.Entry{
		// Log anything that can be marshaled to JSON.
		Payload:  employee,
		Severity: logging.Debug,
	})

}
