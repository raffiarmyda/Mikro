package transactions

import (
	"context"
	"mikro/business/products"
	"mikro/business/users"
	"time"
)

type TransactionUsecase struct {
	Repo           Repository
	UserUsecase    users.Usecase
	ProductUsecase products.Usecase
	Timeout        time.Duration
}

func NewTransactionUsecase(repo Repository, productUsecase products.Usecase, userUsecase users.Usecase, timeout time.Duration) *TransactionUsecase {
	return &TransactionUsecase{
		Repo:           repo,
		UserUsecase:    userUsecase,
		ProductUsecase: productUsecase,
		Timeout:        timeout,
	}
}

func (uc *TransactionUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.TransactionsGetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}
func (uc *TransactionUsecase) GetById(ctx context.Context, domain Domain) (Domain, error) {
	user, err := uc.Repo.TransactionsGetById(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (uc *TransactionUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	user, err := uc.Repo.TransactionsCreate(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *TransactionUsecase) Update(ctx context.Context, domain Domain) (Domain, error) {
	user, err := uc.Repo.TransactionsUpdate(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *TransactionUsecase) Delete(ctx context.Context, domain Domain) error {
	err := uc.Repo.TransactionsDelete(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}
