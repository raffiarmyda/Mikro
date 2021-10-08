package postgres

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mikro/business/products"
	"mikro/repository/databases/records"
)

type ProductRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresProductRepository(conn *gorm.DB) *ProductRepository {
	return &ProductRepository{
		ConnPostgres: conn,
	}
}

func (repo *ProductRepository) ProductsGetById(ctx context.Context, domain products.Domain) (products.Domain, error) {
	var product records.Products
	err := repo.ConnPostgres.Joins("Seller").Find(&product, "products.id = ?", domain.ID)
	if err.Error != nil {
		return products.Domain{}, err.Error
	}
	return product.ProductsToDomain(), nil
}

func (repo *ProductRepository) ProductsCreate(ctx context.Context, domain products.Domain) (products.Domain, error) {
	product := records.ProductsFromDomain(domain)
	err := repo.ConnPostgres.Create(&product)
	if err.Error != nil {
		return products.Domain{}, err.Error
	}
	return product.ProductsToDomain(), nil
}

func (repo *ProductRepository) ProductsGetAll(ctx context.Context) ([]products.Domain, error) {
	var data []records.Products
	err := repo.ConnPostgres.Find(&data)
	if err.Error != nil {
		return []products.Domain{}, err.Error
	}
	return records.ProductsToListDomain(data), nil
}

func (repo *ProductRepository) ProductsUpdate(ctx context.Context, domain products.Domain) (products.Domain, error) {
	data := records.ProductsFromDomain(domain)
	data.Name = domain.Name
	data.Price = domain.Price

	if repo.ConnPostgres.Save(&data).Error != nil {
		return products.Domain{}, errors.New("bad requests")
	}
	return data.ProductsToDomain(), nil
}

func (repo *ProductRepository) ProductsDelete(ctx context.Context, domain products.Domain) error {
	product := records.Products{}
	err := repo.ConnPostgres.Delete(&product, domain.ID)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
