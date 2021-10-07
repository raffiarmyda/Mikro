package records

import (
	"gorm.io/gorm"
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
