package handlers

import (
	"employee-api/data"
	"encoding/json"
	"net/http"

	"cloud.google.com/go/logging"
)

// swagger:route GET /mgrs employees listManagers
// Return a list of employees who are managers from the database
// responses:
//	200: managersResponse

// ReturnAllMgrs handles GET requests and returns all current employees who are managers
func (e *Employees) ReturnAllMgrs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	e.l.StandardLogger(logging.Info).Println("Endpoint Hit: returnAllMgrs")
	var mgrs []data.Employee
	var tempEmp data.Employee
	for index, item := range data.Emps {
		if item.Manager {
			tempEmp = data.Emps[index]
			mgrs = append(mgrs, tempEmp)

		}
	}
	e.l.StandardLogger(logging.Debug).Printf("mgr:%v", mgrs)
	json.NewEncoder(w).Encode(mgrs)
}
