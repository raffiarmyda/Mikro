package transactions

import (
	"context"
)

type Domain struct {
	ID        int
	ProductId int
	BuyerId   int
	Product   Products `json:"product"`
	Buyer     Buyers   `json:"buyer"`
}

type Buyers struct {
	ID          int
	Name        string
	StoreName   string
	City        string
	Phone       string
	Username    string
	BankAccount string
	NoAccount   string
	Password    string
}

type Products struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Users struct {
	Name        string `json:"name"`
	StoreName   string `json:"store_name"`
	City        string `json:"city"`
	Phone       string `json:"phone"`
	BankAccount string `json:"bank_account"`
	NoAccount   string `json:"no_account"`
}

type Usecase interface {
	GetAll(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, domain Domain) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, domain Domain) (Domain, error)
	Delete(ctx context.Context, domain Domain) error
}

type Repository interface {
	TransactionsGetAll(ctx context.Context) ([]Domain, error)
	TransactionsGetById(ctx context.Context, domain Domain) (Domain, error)
	TransactionsCreate(ctx context.Context, domain Domain) (Domain, error)
	TransactionsUpdate(ctx context.Context, domain Domain) (Domain, error)
	TransactionsDelete(ctx context.Context, domain Domain) error
}
