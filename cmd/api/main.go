package main

import (
	"carrierCheck/internal/domain/carrier"
	"carrierCheck/internal/domain/clients"
	"carrierCheck/internal/domain/orders"
	"carrierCheck/internal/domain/products"
	"carrierCheck/internal/endpoints"
	"carrierCheck/internal/infra/database"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

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
	ordersService := orders.OrdersServiceImp{
		Repository: &database.OrdersRepository{
			Db: connection,
		},
	}

	handler := endpoints.Handler{
		CarrierService: &carrierService,
		ClientsService: &clientsService,
		ProductsService: &productsService,
		OrdersService: &ordersService,
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

	router.Post("/orders", endpoints.HandlerError(handler.CreateOrder))
	router.Get("/orders/{id}", endpoints.HandlerError(handler.GetByIdOrder))
	router.Get("/orders", endpoints.HandlerError(handler.GetAllOrder))
	router.Patch("/orders/{id}", endpoints.HandlerError(handler.UpdateOrders))
	router.Patch("/orders/carrier/{id}", endpoints.HandlerError(handler.UpdateCarrierOrder))
	router.Delete("/orders/{id}", endpoints.HandlerError(handler.DeleteOrder))

	http.ListenAndServe(":3000", router)
}