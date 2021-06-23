package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route PUT /emps/{id} employees updateEmployee
// Update a employee details
//
// responses:
//	200: employeeResponse

// Update handles PUT requests to update employee
func UpdateSingleEmp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var tempEmp data.Employee
	for index, item := range data.Emps {
		if strconv.Itoa(item.ID) == vars["id"] {
			_ = json.NewDecoder(r.Body).Decode(&tempEmp)
			tempEmp.ID = index
			data.Emps[index] = tempEmp
			json.NewEncoder(w).Encode(data.Emps[index])
			return
		}
	}
	json.NewEncoder(w).Encode(&data.Employee{})
}
