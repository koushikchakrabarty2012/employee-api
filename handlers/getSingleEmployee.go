package handlers

import (
	"employee-api/data"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//Return single Emp
func ReturnSingleEmp(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("Endpoint Hit: returnSingleEmp")
	w.Header().Set("Content-Type", "application/json")
	key := ps.ByName("id")
	fmt.Print(key)
	for _, item := range data.Emps {
		if strconv.Itoa(item.ID) == key {
			json.NewEncoder(w).Encode(item)
		}
	}
}
