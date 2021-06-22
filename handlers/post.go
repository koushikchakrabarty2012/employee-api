package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Adds a new Emp
func CreateSingleEmp(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var employee data.Employee
	_ = json.NewDecoder(r.Body).Decode(&employee)
	employee.ID = data.Count
	data.Count++
	data.Emps = append(data.Emps, employee)
	json.NewEncoder(w).Encode(employee)

}
