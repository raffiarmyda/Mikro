package postgres

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mikro/business/users"
	"mikro/repository/databases/records"
)

type UserRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		ConnPostgres: conn,
	}
}

func (repo *UserRepository) UsersGetById(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user records.Users
	err := repo.ConnPostgres.Find(&user, "id = ?", domain.ID)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.UsersToDomain(), nil
}

func (repo *UserRepository) UsersGetByUsername(ctx context.Context, domain users.Domain) (users.Domain, error) {
	var user records.Users
	err := repo.ConnPostgres.Find(&user, "username = ?", domain.Username)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.UsersToDomain(), nil
}

func (repo *UserRepository) UsersCreate(ctx context.Context, domain users.Domain) (users.Domain, error) {
	user := records.UsersFromDomain(domain)
	err := repo.ConnPostgres.Create(&user)
	if err.Error != nil {
		return users.Domain{}, err.Error
	}
	return user.UsersToDomain(), nil
}

func (repo *UserRepository) UsersGetAll(ctx context.Context) ([]users.Domain, error) {
	var data []records.Users
	err := repo.ConnPostgres.Find(&data)
	if err.Error != nil {
		return []users.Domain{}, err.Error
	}
	return records.UsersToListDomain(data), nil
}

func (repo *UserRepository) UsersUpdate(ctx context.Context, domain users.Domain) (users.Domain, error) {
	data := records.UsersFromDomain(domain)
	data.Name = domain.Name
	data.StoreName = domain.StoreName
	data.Password = domain.Password
	data.Phone = domain.Phone
	data.City = domain.City
	data.IsAdmin = domain.IsAdmin
	data.Username = domain.Username

	if repo.ConnPostgres.Save(&data).Error != nil {
		return users.Domain{}, errors.New("bad requests")
	}
	return data.UsersToDomain(), nil
}

func (repo *UserRepository) UsersDelete(ctx context.Context, domain users.Domain) error {
	user := records.Users{}
	err := repo.ConnPostgres.Delete(&user, domain.ID)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
