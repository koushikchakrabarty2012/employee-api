package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Updates the Employee with some ID
func UpdateSingleEmp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	key := ps.ByName("id")
	var tempEmp data.Employee
	for index, item := range data.Emps {
		if strconv.Itoa(item.ID) == key {
			_ = json.NewDecoder(r.Body).Decode(&tempEmp)
			tempEmp.ID = index
			data.Emps[index] = tempEmp
			json.NewEncoder(w).Encode(data.Emps[index])
			return
		}
	}
	json.NewEncoder(w).Encode(&data.Employee{})
}
