package users

import (
	"context"
	"mikro/app/middlewares"
	businesses "mikro/business"
	"time"
)

type UserUsecase struct {
	Repo    Repository
	Timeout time.Duration
	jwtAuth *middlewares.ConfigJWT
}

func NewUserUsecase(repo Repository, timeout time.Duration, jwtAuth *middlewares.ConfigJWT) *UserUsecase {
	return &UserUsecase{
		Repo:    repo,
		Timeout: timeout,
		jwtAuth: jwtAuth,
	}
}

func (uc *UserUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}
func (uc *UserUsecase) GetById(ctx context.Context, domain Domain) (Domain, error) {
	user, err := uc.Repo.GetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}
func (uc *UserUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, businesses.ErrEmailRequired
	}
}
func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, string, error) {

}
func (uc *UserUsecase) Update(ctx context.Context, domain Domain) (Domain, error) {

}
func (uc *UserUsecase) Delete(ctx context.Context, domain Domain) error {

}
