package products

import (
	"context"
	"mikro/business/users"
	"time"
)

type ProductUsecase struct {
	Repo        Repository
	UserUsecase users.Usecase
	Timeout     time.Duration
}

func NewProductUsecase(repo Repository, userUsecase users.Usecase, timeout time.Duration) *ProductUsecase {
	return &ProductUsecase{
		Repo:        repo,
		UserUsecase: userUsecase,
		Timeout:     timeout,
	}
}

func (uc *ProductUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.ProductsGetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}
func (uc *ProductUsecase) GetById(ctx context.Context, domain Domain) (Domain, error) {
	user, err := uc.Repo.ProductsGetById(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (uc *ProductUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	user, err := uc.Repo.ProductsCreate(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}

func (uc *ProductUsecase) Update(ctx context.Context, domain Domain) (Domain, error) {
	user, err := uc.Repo.ProductsUpdate(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *ProductUsecase) Delete(ctx context.Context, domain Domain) error {
	err := uc.Repo.ProductsDelete(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}
