package main

import (
	"fmt"
	"net/http"

	"github.com/ctrlmaniac/gocrm/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")

	fmt.Println("Server is running at http://localhost:3000/")
	http.ListenAndServe(":3000", router)
}
