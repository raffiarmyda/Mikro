package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	StoreName string
	City      string
	Phone     string
	Username  string
	IsAdmin   bool
	Password  string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, domain Domain) error
}

type Repository interface {
	UsersGetByUsername(ctx context.Context, domain Domain) (Domain, error)
	UsersGetAll(ctx context.Context) ([]Domain, error)
	UsersGetById(ctx context.Context, domain Domain) (Domain, error)
	UsersCreate(ctx context.Context, domain Domain) (Domain, error)
	UsersUpdate(ctx context.Context, domain Domain) (Domain, error)
	UsersDelete(ctx context.Context, domain Domain) error
}
