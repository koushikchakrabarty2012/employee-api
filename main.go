package main

import (
	"bytes"
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
	Address   Address
}

//Emps defines slice of employees
var Emps []Employee

func main() {
	baseurl := "https://employee-api-sjcc6yr6ma-ew.a.run.app/emps"
	getAllEmployees(baseurl)
	//createSingleEmployee(baseurl)
	//getSingleEmployee(baseurl, "2")
	//updateSingleEmployee(baseurl, "3")
}

func getAllEmployees(baseurl string) {
	fmt.Println("Printing all employees which are present")
	resp, err := http.Get(baseurl)
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

func getSingleEmployee(baseurl string, id string) {
	fmt.Printf("Printing single employee whose id %s has been paassed", id)
	resp, err := http.Get(baseurl + "/" + id)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
	var emp Employee
	json.Unmarshal(bodyBytes, &emp)
	fmt.Printf("API Response as struct %+v\n", emp)

}

func createSingleEmployee(baseurl string) {
	fmt.Println("Performing Http Post...")
	c_emp := Employee{FirstName: "firstName2", LastName: "lastName2", Manager: true, Address: Address{City: "city2", State: "state2", Pin: "pin2"}}
	jsonReq, err := json.Marshal(c_emp)
	request, _ := http.NewRequest("POST", baseurl, bytes.NewBuffer(jsonReq))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	//If we choose to send JSON data, we need to marshal it into an array of bytes. We also need to define in the header that we are sending JSON data, just in case the recipient doesn’t make assumptions.
	//response, err = http.Post(baseurl, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
	var emp Employee
	json.Unmarshal(bodyBytes, &emp)
	fmt.Printf("API Response as struct %+v\n", emp)

}

func updateSingleEmployee(baseurl string, id string) {
	fmt.Println("Performing Http Put...")
	c_emp := Employee{FirstName: "firstName3", LastName: "lastName3", Manager: true, Address: Address{City: "city3", State: "state3", Pin: "pin3"}}
	jsonReq, err := json.Marshal(c_emp)
	request, _ := http.NewRequest("PUT", baseurl+"/"+id, bytes.NewBuffer(jsonReq))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)
	var emp Employee
	json.Unmarshal(bodyBytes, &emp)
	fmt.Printf("API Response as struct %+v\n", emp)

}
