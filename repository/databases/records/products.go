package records

import (
	"gorm.io/gorm"
	"mikro/business/products"
	"time"
)

type Products struct {
	ID        int `gorm:"primaryKey"`
	SellerId  int
	Seller    Users `gorm:"foreignKey:seller_id"`
	Name      string
	Price     float64
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *Products) ProductsToDomain() products.Domain {
	return products.Domain{
		ID:        p.ID,
		SellerId:  p.SellerId,
		Seller:    FromProductToSeller(p),
		Name:      p.Name,
		Price:     p.Price,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func ProductsToListDomain(data []Products) []products.Domain {
	var list []products.Domain
	for _, v := range data {
		list = append(list, v.ProductsToDomain())
	}
	return list
}

func ProductsFromDomain(domain products.Domain) Products {
	return Products{
		ID:       domain.ID,
		Name:     domain.Name,
		Price:    domain.Price,
		SellerId: domain.SellerId,
	}
}

func FromProductToSeller(p *Products) products.Seller {
	return products.Seller{
		StoreName: p.Seller.StoreName,
		Name:      p.Seller.Name,
		City:      p.Seller.City,
	}
}
