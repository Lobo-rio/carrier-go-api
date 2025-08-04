package main

import (
	"carrierCheck/internal/domain/carrier"
	"carrierCheck/internal/domain/clients"
	"carrierCheck/internal/domain/products"
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
	productsService := products.ProductsServiceImp{
		Repository: &database.ProductsRepository{
			Db: connection,
		},
	}

	handler := endpoints.Handler{
		CarrierService: &carrierService,
		ClientsService: &clientsService,
		ProductsService: &productsService,
	}

	router.Post("/carriers", endpoints.HandlerError(handler.CreateCarrier))
	router.Get("/carriers/{id}", endpoints.HandlerError(handler.GetByIdCarrier))
	router.Get("/carriers", endpoints.HandlerError(handler.GetAllCarrier))
	router.Patch("/carriers/{id}", endpoints.HandlerError(handler.UpdateCarrier))
	router.Delete("/carriers/{id}", endpoints.HandlerError(handler.DeleteCarrier))

	router.Post("/clients", endpoints.HandlerError(handler.CreateClient))
	router.Get("/clients/{id}", endpoints.HandlerError(handler.GetByIdClient))
	router.Get("/clients", endpoints.HandlerError(handler.GetAllClient))
	router.Patch("/clients/{id}", endpoints.HandlerError(handler.UpdateClient))
	router.Delete("/clients/{id}", endpoints.HandlerError(handler.DeleteClient))
	
    router.Post("/products", endpoints.HandlerError(handler.CreateProduct))
	router.Get("/products/{id}", endpoints.HandlerError(handler.GetByIdProduct))
	router.Get("/products", endpoints.HandlerError(handler.GetAllProduct))
	router.Patch("/products/{id}", endpoints.HandlerError(handler.UpdateProduct))
	router.Delete("/products/{id}", endpoints.HandlerError(handler.DeleteProduct))

	http.ListenAndServe(":3000", router)
}