package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"
)

// swagger:route POST /emps employees createEmployee
// Create a new Employee
//
// responses:
//	200: employeeResponse

// Create handles POST requests to add new employees
func CreateSingleEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee data.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = data.Count
	data.Count++
	data.Emps = append(data.Emps, employee)
	json.NewEncoder(w).Encode(employee)

}
