package requests

import "mikro/business/users"

type UsersRequest struct {
	ID        int    `json:"id"`
	StoreName string `json:"store_name"`
	City      string `json:"city"`
	Phone     string `json:"phone"`
	IsAdmin   bool   `json:"is_admin"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func (ur *UsersRequest) ToDomainUser() users.Domain {
	return users.Domain{
		ID:        ur.ID,
		StoreName: ur.StoreName,
		City:      ur.City,
		Phone:     ur.Phone,
		IsAdmin:   ur.IsAdmin,
		Username:  ur.Username,
		Password:  ur.Password,
	}
}
