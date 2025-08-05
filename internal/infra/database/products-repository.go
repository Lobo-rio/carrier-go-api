package database

import (
	"carrierCheck/internal/domain/products"

	"gorm.io/gorm"
)

type ProductsRepository struct{
	Db *gorm.DB
}

func (c *ProductsRepository) Save(product *products.Product) error {
	tx := c.Db.Create(product)
	return tx.Error
}

func (c *ProductsRepository) GetAll() ([]products.Product, error) {
	var products []products.Product
	tx := c.Db.Find(&products)
	return products, tx.Error
}

func (c *ProductsRepository) GetById(id string) (*products.Product, error) {
	var product products.Product
	tx := c.Db.First(&product, "id = ?", id)
	return &product, tx.Error
}

func (c *ProductsRepository) Update(product *products.Product) error {
	tx := c.Db.Save(product)
    return tx.Error
}

func (c *ProductsRepository) Delete(product *products.Product) error {
	tx := c.Db.Delete(product)
	return tx.Error
}