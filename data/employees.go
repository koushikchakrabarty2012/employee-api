package data

import (
	"fmt"
	"log"
)

// ErrEmployeeNotFound is an error raised when an employee can not be found in the database
var ErrEmployeeNotFound = fmt.Errorf("Employee not found")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// Used for providing a unique ID to each book
var Count int

// Address defines the structure for an  employee address
// swagger:model
type Address struct {
	City  string `json:"residingCity"`
	State string `json:"stateCode"`
	Pin   string `json:"pinCode"`
}

// Employee defines the structure for an  employee
// swagger:model
type Employee struct {
	// the id for the employee
	//
	// required: false
	// min: 1
	ID        int    `json:"id"` // Unique identifier for the employee
	FirstName string `json:"empFirstName"`
	LastName  string `json:"empLastName"`
	Manager   bool   `json:"isManager"`
	Address   *Address
}

//Emps defines slice of employees
var Emps []Employee

// findIndex finds the index of an employee in the database
// returns -1 when no employee can be found
func findIndexByEmployeeID(id int) int {
	for i, p := range Emps {
		if p.ID == id {
			return i
		}
	}

	return -1
}

// GetEmployeeByID returns a single employee which matches the id from the  database.

// If an employee is not found this function returns a EmployeeNotFound error
func GetEmployeeByID(id int) (*Employee, error) {
	i := findIndexByEmployeeID(id)
	log.Println("current value if index", i)
	if i == -1 {
		log.Println("inside if before throwing error")
		return nil, ErrEmployeeNotFound
	}

	return (&Emps[i]), nil
}
