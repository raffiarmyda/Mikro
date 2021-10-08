package records

import (
	"gorm.io/gorm"
	"mikro/business/users"
	"time"
)

type Users struct {
	ID          int `gorm:"primaryKey"`
	Name        string
	StoreName   string
	City        string
	Phone       string
	BankAccount string
	NoAccount   string
	IsAdmin     bool
	Username    string
	Password    string
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (user *Users) UsersToDomain() users.Domain {
	return users.Domain{
		ID:          user.ID,
		Name:        user.Name,
		City:        user.City,
		Phone:       user.Phone,
		StoreName:   user.StoreName,
		BankAccount: user.BankAccount,
		Username:    user.Username,
		IsAdmin:     user.IsAdmin,
		Password:    user.Password,
		NoAccount:   user.NoAccount,
		Token:       user.StoreName,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}

func UsersToListDomain(data []Users) []users.Domain {
	var list []users.Domain
	for _, v := range data {
		list = append(list, v.UsersToDomain())
	}
	return list
}

func UsersFromDomain(domain users.Domain) Users {
	return Users{
		ID:          domain.ID,
		Name:        domain.Name,
		City:        domain.City,
		BankAccount: domain.BankAccount,
		Phone:       domain.Phone,
		NoAccount:   domain.NoAccount,
		StoreName:   domain.StoreName,
		Username:    domain.Username,
		IsAdmin:     domain.IsAdmin,
		Password:    domain.Password,
	}
}
