package requests

import (
	"mikro/business/products"
)

type ProductsRequest struct {
	SellerId int     `json:"seller_id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
}

func (ur *ProductsRequest) ToDomainProduct() products.Domain {
	return products.Domain{
		SellerId: ur.SellerId,
		Name:     ur.Name,
		Price:    ur.Price,
	}
}
