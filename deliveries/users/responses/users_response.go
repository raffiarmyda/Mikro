package responses

import (
	"mikro/business/users"
	"time"
)

type UsersResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	StoreName string    `json:"store_name"`
	City      string    `json:"city"`
	Phone     string    `json:"phone"`
	IsAdmin   bool      `json:"is_admin"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromUsersDomainToLogin(domain users.Domain, token string) LoginResponse {
	response := UsersResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		StoreName: domain.StoreName,
		City:      domain.City,
		Phone:     domain.Phone,
		IsAdmin:   domain.IsAdmin,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
	loginResponse := LoginResponse{}
	loginResponse.SessionToken = token
	loginResponse.User = response
	return loginResponse
}

type LoginResponse struct {
	SessionToken string      `json:"session_token"`
	User         interface{} `json:"user"`
}

func FromUsersDomain(domain users.Domain) UsersResponse {
	return UsersResponse{
		ID:        domain.ID,
		Name:      domain.Name,
		StoreName: domain.StoreName,
		City:      domain.City,
		Phone:     domain.Phone,
		IsAdmin:   domain.IsAdmin,
		Username:  domain.Username,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromUsersListDomain(domain []users.Domain) []UsersResponse {
	var list []UsersResponse
	for _, v := range domain {
		list = append(list, FromUsersDomain(v))
	}
	return list
}
