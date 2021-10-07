package records

import "time"

type Products struct {
	ID        int `gorm:"primaryKey"`
	SellerId  int
	Seller    Users `gorm:"foreignKey:seller_id"`
	Addres    string
	Phone     string
	Username  string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
