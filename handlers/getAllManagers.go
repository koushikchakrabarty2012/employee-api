package handlers

import (
	"employee-api/data"
	"encoding/json"
	"fmt"
	"net/http"
)

// swagger:route GET /mgrs employees listManagers
// Return a list of employees who are managers from the database
// responses:
//	200: managersResponse

// ReturnAllMgrs handles GET requests and returns all current employees who are managers
func ReturnAllMgrs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: returnAllMgrs")
	var mgrs []data.Employee
	var tempEmp data.Employee
	for index, item := range data.Emps {
		if item.Manager {
			tempEmp = data.Emps[index]
			mgrs = append(mgrs, tempEmp)

		}
	}
	fmt.Printf("mgr:%v", mgrs)
	json.NewEncoder(w).Encode(mgrs)
}
