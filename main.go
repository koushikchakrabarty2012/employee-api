package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Address defines the structure for an  employee address

type Address struct {
	City  string `json:"residingCity"`
	State string `json:"stateCode"`
	Pin   string `json:"pinCode"`
}

// Employee defines the structure for an  employee

type Employee struct {
	ID        int    `json:"id"`
	FirstName string `json:"empFirstName"`
	LastName  string `json:"empLastName"`
	Manager   bool   `json:"isManager"`
	Address   *Address
}

//Emps defines slice of employees
var Emps []Employee

func main() {
	getAllEmployees()
}

func getAllEmployees() {
	fmt.Println("Printing all employees which are present")
	resp, err := http.Get("https://employee-api-sjcc6yr6ma-ew.a.run.app/emps")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	json.Unmarshal(bodyBytes, &Emps)
	fmt.Printf("API Response as struct %+v\n", Emps)

}
