package handlers

import (
	"employee-api/data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:route GET /emps/{id} employees listSingleEmployee
// Return single employee from the database
// responses:
//	200: employeeResponse
// ReturnSingleEmp handles GET requests
func ReturnSingleEmp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleEmp")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Print(key)
	for _, item := range data.Emps {
		if strconv.Itoa(item.ID) == key {
			json.NewEncoder(w).Encode(item)
		}
	}
}
