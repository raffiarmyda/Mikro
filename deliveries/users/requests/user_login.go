package requests

import (
	"mikro/business/users"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (login *UserLogin) ToDomain() users.Domain {
	return users.Domain{
		Username: login.Username,
		Password: login.Password,
	}
}
