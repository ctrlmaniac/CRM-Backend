package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ctrlmaniac/gocrm/persistance"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(persistance.Data)
}
