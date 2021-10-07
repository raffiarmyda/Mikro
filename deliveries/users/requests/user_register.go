package requests

import (
	"mikro/business/users"
)

type UserRegister struct {
	Name      string `json:"name"`
	StoreName string `json:"store_name"`
	City      string `json:"city"`
	Phone     string `json:"phone"`
	IsAdmin   bool   `json:"is_admin"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (ur UserRegister) ToDomain() users.Domain {
	return users.Domain{
		Name:      ur.Name,
		StoreName: ur.StoreName,
		City:      ur.City,
		Phone:     ur.Phone,
		Username:  ur.Username,
		IsAdmin:   ur.IsAdmin,
		Password:  ur.Password,
	}
}
