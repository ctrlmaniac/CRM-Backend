package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ctrlmaniac/gocrm/models"
	"github.com/ctrlmaniac/gocrm/persistance"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(persistance.Data)
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newEntry map[string]string

	reqBody, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(reqBody, &newEntry)

	if _, hasID := newEntry["ID"]; hasID {
		// if request has ID it will return an error
		w.WriteHeader(http.StatusConflict)
	} else {
		contacted, _ := strconv.ParseBool(newEntry["Contacted"])

		persistance.Data = append(persistance.Data, models.Customer{
			ID:        uuid.New().String(),
			Name:      newEntry["Name"],
			Role:      newEntry["Role"],
			Email:     newEntry["Email"],
			Phone:     newEntry["Phone"],
			Contacted: contacted,
		})

		json.NewEncoder(w).Encode(persistance.Data)
	}
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	for _, c := range persistance.Data {
		if c.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(c)
			break
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func UpdateCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	for i, c := range persistance.Data {
		if c.ID == id {
			// If customer is found then update
			var rawData map[string]string

			reqBody, _ := ioutil.ReadAll(r.Body)

			json.Unmarshal(reqBody, &rawData)

			contacted, _ := strconv.ParseBool(rawData["Contacted"])

			updatedCustomer := models.Customer{
				ID:        id,
				Name:      rawData["Name"],
				Role:      rawData["Role"],
				Email:     rawData["Email"],
				Phone:     rawData["Phone"],
				Contacted: contacted,
			}

			persistance.Data[i] = updatedCustomer

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedCustomer)
			break
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func DeleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	for i, c := range persistance.Data {
		if c.ID == id {
			persistance.Data = append(persistance.Data[:i], persistance.Data[i+1:]...)

			w.WriteHeader(http.StatusOK)
			break
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
