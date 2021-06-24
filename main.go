package main

import (
	"fmt"
	"log"
	"net/http"

	"employee-api/data"
	"employee-api/handlers"

	"github.com/go-openapi/runtime/middleware"

	"github.com/gorilla/mux"
)

func main() {
	//Init router
	myRouter := mux.NewRouter()
	getRouter := myRouter.Methods(http.MethodGet).Subrouter()
	putRouter := myRouter.Methods(http.MethodPut).Subrouter()
	postRouter := myRouter.Methods(http.MethodPost).Subrouter()
	deleteRouter := myRouter.Methods(http.MethodDelete).Subrouter()
	//generate mock data
	data.Emps = append(data.Emps, data.Employee{ID: data.Count, FirstName: "firstname0", LastName: "lastname0", Manager: true, Address: &data.Address{City: "city0", State: "state0", Pin: "pin0"}})
	data.Count++
	data.Emps = append(data.Emps, data.Employee{ID: data.Count, FirstName: "firstname1", LastName: "lastname1", Manager: false, Address: &data.Address{City: "city1", State: "state1", Pin: "pin1"}})
	data.Count++
	fmt.Printf("%+v", data.Emps)

	//Route Handler Endpoints
	//getRouter.GET("/", handlers.HomePage)
	getRouter.HandleFunc("/emps", handlers.ReturnAllEmps)
	getRouter.HandleFunc("/emps/{id}", handlers.ReturnSingleEmp)
	postRouter.HandleFunc("/emps", handlers.CreateSingleEmp)
	putRouter.HandleFunc("/emps/{id}", handlers.UpdateSingleEmp)
	deleteRouter.HandleFunc("/emps/{id}", handlers.DeleteSingleEmp)
	getRouter.HandleFunc("/mgrs", handlers.ReturnAllMgrs)

	//handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	//starting server
	log.Println("Starting server on PORT:8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
