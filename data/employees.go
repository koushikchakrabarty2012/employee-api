package data

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
	ID        int    `json:"id"`
	FirstName string `json:"empFirstName"`
	LastName  string `json:"empLastName"`
	Manager   bool   `json:"isManager"`
	Address   *Address
}

//Emps defines slice of employees
var Emps []Employee
