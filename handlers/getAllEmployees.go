package handlers

import (
	"employee-api/data"
	"encoding/json"
	"fmt"
	"net/http"
)

// swagger:route GET /emps employees listEmployees
// Return a list of employees from the database
// responses:
//	200: employeesResponse

// ReturnAllEmps handles GET requests and returns all current employees
func ReturnAllEmps(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: returnAllEmps")
	fmt.Println(data.Emps)
	json.NewEncoder(w).Encode(data.Emps)

}
