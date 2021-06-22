package handlers

import (
	"employee-api/data"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//Return all Emps
func ReturnAllEmps(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Endpoint Hit: returnAllEmps")
	fmt.Println(data.Emps)
	json.NewEncoder(w).Encode(data.Emps)

}
