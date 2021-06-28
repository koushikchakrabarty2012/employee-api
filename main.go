package main

import (
	"fmt"
	"log"
	"net/http"

	"context"
	"employee-api/data"
	"employee-api/handlers"

	"cloud.google.com/go/logging"
	"github.com/go-openapi/runtime/middleware"

	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "ingka-fulfilment-ordermgt-dev"

	// Creates a client.
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	logName := "employee-api-log"

	// create the handlers
	eh := handlers.NewEmployees(client.Logger(logName))
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
	getRouter.HandleFunc("/emps", eh.ReturnAllEmps)
	getRouter.HandleFunc("/emps/{id}", eh.ReturnSingleEmp)
	postRouter.HandleFunc("/emps", eh.CreateSingleEmp)
	putRouter.HandleFunc("/emps/{id}", eh.UpdateSingleEmp)
	deleteRouter.HandleFunc("/emps/{id}", eh.DeleteSingleEmp)
	getRouter.HandleFunc("/mgrs", eh.ReturnAllMgrs)

	//handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	//starting server
	client.Logger(logName).StandardLogger(logging.Info).Println("Starting server on PORT:8081")
	client.Logger(logName).StandardLogger(logging.Error).Fatal(http.ListenAndServe(":8081", myRouter))
}
