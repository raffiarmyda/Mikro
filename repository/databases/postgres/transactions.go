package postgres

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"mikro/business/transactions"
	"mikro/repository/databases/records"
)

type TransactionRepository struct {
	ConnPostgres *gorm.DB
}

func NewPostgresTransactionRepository(conn *gorm.DB) *TransactionRepository {
	return &TransactionRepository{
		ConnPostgres: conn,
	}
}

func (repo *TransactionRepository) TransactionsGetById(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	var product records.Transactions
	err := repo.ConnPostgres.Find(&product, "id = ?", domain.ID)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return product.TransactionsToDomain(), nil
}

func (repo *TransactionRepository) TransactionsCreate(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	product := records.TransactionsFromDomain(domain)
	err := repo.ConnPostgres.Create(&product)
	if err.Error != nil {
		return transactions.Domain{}, err.Error
	}
	return product.TransactionsToDomain(), nil
}

func (repo *TransactionRepository) TransactionsGetAll(ctx context.Context) ([]transactions.Domain, error) {
	var data []records.Transactions
	err := repo.ConnPostgres.Preload("Product").Joins("Buyer").Find(&data)
	if err.Error != nil {
		return []transactions.Domain{}, err.Error
	}
	return records.TransactionsToListDomain(data), nil
}

func (repo *TransactionRepository) TransactionsUpdate(ctx context.Context, domain transactions.Domain) (transactions.Domain, error) {
	data := records.TransactionsFromDomain(domain)
	data.BuyerId = domain.BuyerId
	data.ProductId = domain.ProductId

	if repo.ConnPostgres.Save(&data).Error != nil {
		return transactions.Domain{}, errors.New("bad requests")
	}
	return data.TransactionsToDomain(), nil
}

func (repo *TransactionRepository) TransactionsDelete(ctx context.Context, domain transactions.Domain) error {
	product := records.Transactions{}
	err := repo.ConnPostgres.Delete(&product, domain.ID)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
