package products

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	SellerId  int
	Seller    Seller
	Name      string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Seller struct {
	StoreName string `json:"store_name"`
	Name      string `json:"name"`
	City      string `json:"city"`
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, domain Domain) error
}

type Repository interface {
	ProductsGetAll(ctx context.Context) ([]Domain, error)
	ProductsGetById(ctx context.Context, domain Domain) (Domain, error)
	ProductsCreate(ctx context.Context, domain Domain) (Domain, error)
	ProductsUpdate(ctx context.Context, domain Domain) (Domain, error)
	ProductsDelete(ctx context.Context, domain Domain) error
}
