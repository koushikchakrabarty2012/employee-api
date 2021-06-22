package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Deletes the Employee with some ID
func DeleteSingleEmp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	key := ps.ByName("id")
	var tempEmp data.Employee
	for index, item := range data.Emps {
		if strconv.Itoa(item.ID) == key {
			tempEmp = data.Emps[index]
			data.Emps = append(data.Emps[:index], data.Emps[index+1:]...)
			_ = json.NewEncoder(w).Encode(tempEmp)
			return
		}
	}
	json.NewEncoder(w).Encode(&data.Employee{})
}
