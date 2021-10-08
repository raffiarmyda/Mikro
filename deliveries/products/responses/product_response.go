package responses

import (
	"mikro/business/products"
	"time"
)

type ProductResponse struct {
	SellerId  int         `json:"seller_id"`
	Seller    interface{} `json:"seller"`
	Name      string      `json:"name"`
	Price     float64     `json:"price"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func FromProductDomain(domain products.Domain) ProductResponse {
	return ProductResponse{
		SellerId:  domain.SellerId,
		Seller:    domain.Seller,
		Name:      domain.Name,
		Price:     domain.Price,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromProductListDomain(listDomain []products.Domain) []ProductResponse {
	var list []ProductResponse
	for _, v := range listDomain {
		list = append(list, FromProductDomain(v))
	}
	return list
}
