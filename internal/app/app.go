package app

import (
	"context"
	"ddd-arch/internal/driver"
	"ddd-arch/internal/services"
	"fmt"

	"github.com/go-chi/chi/v5"
)

type Application struct {
	tavernDriver *driver.TavernHTTPDriver
	router       chi.Router
}

func NewApplication() *Application {
	return new(Application)
}

func (app *Application) Init(ctx context.Context) error {
	orderService, err := services.NewOrderService(services.WithMemoryCustomerRepository())
	if err != nil {
		return fmt.Errorf("failed to init order service: %w", err)
	}

	tavernService, err := services.NewTavern(services.WithOrderService(orderService))
	if err != nil {
		return fmt.Errorf("failed to init tavern service: %w", err)
	}

	app.tavernDriver = driver.NewTavernHTTPDriver(tavernService)
	app.router = app.initRouter()

	return nil
}

func (app *Application) Run(ctx context.Context) error {
	return nil
}

func (app *Application) initRouter() chi.Router {
	router := chi.NewRouter()
	router.Post("/order", nil)
	return router
}
