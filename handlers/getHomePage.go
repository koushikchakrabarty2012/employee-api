package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//HomePage
func HomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to the HomePage!\n")
	fmt.Println("Endpoint Hit: homePage")
}
