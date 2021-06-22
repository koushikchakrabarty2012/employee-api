package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"employee-api/data"
	"employee-api/handlers"
)

func main() {
	//Init router
	myRouter := httprouter.New()

	//generate mock data
	data.Emps = append(data.Emps, data.Employee{ID: data.Count, FirstName: "firstname0", LastName: "lastname0", Manager: true, Address: &data.Address{City: "city0", State: "state0", Pin: "pin0"}})
	data.Count++
	data.Emps = append(data.Emps, data.Employee{ID: data.Count, FirstName: "firstname1", LastName: "lastname1", Manager: false, Address: &data.Address{City: "city1", State: "state1", Pin: "pin1"}})
	data.Count++
	fmt.Printf("%+v", data.Emps)

	//Route Handler Endpoints
	myRouter.GET("/", handlers.HomePage)
	myRouter.GET("/emps", handlers.ReturnAllEmps)
	myRouter.GET("/emps/:id", handlers.ReturnSingleEmp)
	myRouter.POST("/emps", handlers.CreateSingleEmp)
	myRouter.PUT("/emps/:id", handlers.UpdateSingleEmp)
	myRouter.DELETE("/emps/:id", handlers.DeleteSingleEmp)
	myRouter.GET("/mgrs", handlers.ReturnAllMgrs)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
