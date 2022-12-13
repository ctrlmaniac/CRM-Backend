package main

import (
	"fmt"
	"net/http"

	"github.com/ctrlmaniac/gocrm/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", controllers.Home).Methods("GET")

	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", controllers.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers/{id}", controllers.UpdateCustomerByID).Methods("PUT")
	router.HandleFunc("/customers/{id}", controllers.DeleteCustomerByID).Methods("DELETE")

	fmt.Println("Server is running at http://localhost:3000/")
	http.ListenAndServe(":3000", router)
}
