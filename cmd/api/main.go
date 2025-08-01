package main

import (
	"carrierCheck/internal/domain/carrier"
	"carrierCheck/internal/domain/clients"
	"carrierCheck/internal/endpoints"
	"carrierCheck/internal/infra/database"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	connection := database.Connection()
	carrierService := carrier.CarrierServiceImp{
		Repository: &database.CarrierRepository{
			Db: connection,
		},
	}
	clientsService := clients.ClientsServiceImp{
		Repository: &database.ClientsRepository{
			Db: connection,
		},
	}
	handler := endpoints.Handler{
		CarrierService: &carrierService,
		ClientsService: &clientsService,
	}

	router.Post("/carriers", endpoints.HandlerError(handler.CreateCarrier))
	router.Get("/carriers/{id}", endpoints.HandlerError(handler.GetByIdCarrier))
	router.Get("/carriers", endpoints.HandlerError(handler.GetAllCarrier))
	router.Patch("/carriers/{id}", endpoints.HandlerError(handler.UpdateCarrier))
	router.Delete("/carriers/{id}", endpoints.HandlerError(handler.DeleteCarrier))

	router.Post("/clients", endpoints.HandlerError(handler.CreateClients))
	
	http.ListenAndServe(":3000", router)
}