package database

import (
	"carrierCheck/internal/domain/carrier"
	"carrierCheck/internal/domain/clients"
	"carrierCheck/internal/domain/orders"
	"carrierCheck/internal/domain/products"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
    dsn := "host=" + os.Getenv("DB_HOST") +
	" user=" + os.Getenv("DB_USER") +
	" password=" + os.Getenv("DB_PASSWORD") +
	" dbname=" + os.Getenv("DB_NAME") +
	" port=" + os.Getenv("DB_PORT") +
	" sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&carrier.Carrier{})
	db.AutoMigrate(&carrier.EmailCarrier{})
	db.AutoMigrate(&products.Product{})
	db.AutoMigrate(&clients.Client{})
	db.AutoMigrate(&clients.AddressClients{})
	db.AutoMigrate(&orders.Order{})
	db.AutoMigrate(&orders.OrdersProducts{})

	return db
}