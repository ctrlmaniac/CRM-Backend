package persistance

import (
	"github.com/ctrlmaniac/gocrm/models"
	"github.com/google/uuid"
)

var customers = []models.Customer{
	{
		ID:        uuid.New().String(),
		Name:      "Mario Rossi",
		Role:      "customer",
		Email:     "mario.rossi@email.com",
		Phone:     "1515232345",
		Contacted: true,
	},
	{
		ID:        uuid.New().String(),
		Name:      "Eleonora Gialli",
		Role:      "customer",
		Email:     "eleonora.gialli@email.com",
		Phone:     "5432325151",
		Contacted: true,
	},
	{
		ID:        uuid.New().String(),
		Name:      "Gino Verdi",
		Role:      "customer",
		Email:     "gino.verdi@email.com",
		Phone:     "9090878765",
		Contacted: true,
	},
}
