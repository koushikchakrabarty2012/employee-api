package handlers

import (
	"fmt"
	"net/http"
)

//HomePage
func (e *Employees) HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!\n")
	fmt.Println("Endpoint Hit: homePage")
}
