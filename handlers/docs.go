// Package classification of Employee API
//
// Documentation for Employee API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "employee-api/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// A list of employees
// swagger:response employeesResponse
type employeesResponseWrapper struct {
	// All current employees
	// in: body
	Body []data.Employee
}

// A list of managers
// swagger:response  managersResponse
type managersResponseWrapper struct {
	// All current managers
	// in: body
	Body []data.Employee
}

// swagger:parameters updateEmployee createEmployee
type employeeParamsWrapper struct {
	// Employee data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Employee
}

// swagger:parameters  deleteEmployee listSingleEmployee
type employeeIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
