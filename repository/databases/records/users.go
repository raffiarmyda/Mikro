package records

import "time"

type Users struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Addres    string
	Phone     string
	Username  string
	Password  string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
