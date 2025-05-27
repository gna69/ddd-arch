package driver

import (
	"ddd-arch/internal/services"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type TavernHTTPDriver struct {
	service *services.Tavern
}

func NewTavernHTTPDriver(service *services.Tavern) *TavernHTTPDriver {
	return &TavernHTTPDriver{}
}

func (d *TavernHTTPDriver) CreateOrder(w http.ResponseWriter, r *http.Request) {

	/*
		unmarshal body to dto
		validate input
		convert dto to entities
	*/

	customerID := uuid.New()
	productsIDs := []uuid.UUID{uuid.New(), uuid.New()}
	price, err := d.service.OrderService.CreateOrder(customerID, productsIDs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%f", price)))
}
