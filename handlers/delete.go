package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route DELETE /emps/{id} employees deleteEmployee
// Delete a employee details
//
// responses:
//	200: employeeResponse

// Delete handles DELETE requests and removes employee from the database
func (e *Employees) DeleteSingleEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var tempEmp data.Employee
	for index, item := range data.Emps {
		if strconv.Itoa(item.ID) == vars["id"] {
			tempEmp = data.Emps[index]
			data.Emps = append(data.Emps[:index], data.Emps[index+1:]...)
			_ = json.NewEncoder(w).Encode(tempEmp)
			return
		}
	}
	json.NewEncoder(w).Encode(&data.Employee{})
}
