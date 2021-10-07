package users

import (
	"context"
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Addres    string
	Phone     string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, string, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, domain Domain) error
}

type Repository interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, domain Domain) error
}
