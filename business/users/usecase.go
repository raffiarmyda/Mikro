package users

import (
	"context"
	"log"
	"mikro/app/middlewares"
	businesses "mikro/business"
	"mikro/helpers"
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
	user, err := uc.Repo.UsersGetAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return user, nil
}
func (uc *UserUsecase) GetById(ctx context.Context, domain Domain) (Domain, error) {
	user, err := uc.Repo.UsersGetById(ctx, domain)
	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (uc *UserUsecase) Create(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, businesses.ErrUsernameRequired
	}
	if domain.Password == "" {
		return Domain{}, businesses.ErrPasswordRequired
	}
	check, err := uc.Repo.UsersGetByUsername(ctx, domain)
	if check.ID != 0 {
		return Domain{}, businesses.ErrUsernameHasBeenRegister
	}
	domain.Password, _ = helpers.HashPassword(domain.Password)

	user, err := uc.Repo.UsersCreate(ctx, domain)

	if err != nil {
		return Domain{}, err
	}
	return user, nil
}
func (uc *UserUsecase) Login(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" || domain.Password == "" {
		return Domain{}, businesses.ErrUsernamePasswordNotFound
	}

	user, err := uc.Repo.UsersGetByUsername(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	if !helpers.CheckPassword(domain.Password, user.Password) {
		return Domain{}, businesses.ErrInvalidAuthentication
	}

	//JWT
	token, errToken := uc.jwtAuth.GenerateTokenJWT(user.ID, user.IsAdmin)
	if errToken != nil {
		log.Println(errToken)
	}
	if token == "" {
		return Domain{}, businesses.ErrInvalidAuthentication
	}
	user.Token = token

	return user, nil
}
func (uc *UserUsecase) Update(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Username == "" || domain.Password == "" {
		return Domain{}, businesses.ErrUsernamePasswordNotFound
	}
	if domain.Username != "" {
		data, err := uc.Repo.UsersGetByUsername(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		if data.ID > 0 && data.ID != domain.ID {
			return Domain{}, businesses.ErrUsernameHasBeenRegister
		}
	}
	domain.Password, _ = helpers.HashPassword(domain.Password)
	user, err := uc.Repo.UsersUpdate(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
func (uc *UserUsecase) Delete(ctx context.Context, domain Domain) error {
	err := uc.Repo.UsersDelete(ctx, domain)
	if err != nil {
		return err
	}
	return nil
}
