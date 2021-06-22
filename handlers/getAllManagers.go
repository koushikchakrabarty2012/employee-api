package handlers

import (
	"employee-api/data"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//return all employees who are managers
func ReturnAllMgrs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
