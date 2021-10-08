package transactions_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	businesses "mikro/business"
	_mockProducts "mikro/business/products/mocks"
	"mikro/business/transactions"
	_mockTransactions "mikro/business/transactions/mocks"
	_mockUsers "mikro/business/users/mocks"

	"testing"
	"time"
)

var transactionRepository _mockTransactions.Repository
var userService _mockUsers.Usecase
var productService _mockProducts.Usecase

var transactionService transactions.Usecase

var transactionDomain transactions.Domain
var listTransactionDomain []transactions.Domain

func setup() {
	transactionService = transactions.NewTransactionUsecase(&transactionRepository, &productService, &userService, time.Second*10)
	transactionDomain = transactions.Domain{
		ID:        1,
		ProductId: 1,
		BuyerId:   1,
	}
	listTransactionDomain = append(listTransactionDomain, transactionDomain)

}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsGetAll",
			mock.Anything).Return(listTransactionDomain, nil).Once()
		data, err := transactionService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listTransactionDomain))
		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2| UsersGetAll - Error", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsGetAll",
			mock.Anything).Return([]transactions.Domain{}, businesses.ErrForTest).Once()
		data, err := transactionService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []transactions.Domain{})
		transactionRepository.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | TransactionsGetById - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsGetById",
			mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()
		data, err := transactionService.GetById(context.Background(), transactionDomain)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetById - Error", func(t *testing.T) {
		setup()
		transactionDomain.ID = 0
		transactionRepository.On("TransactionsGetById",
			mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()
		data, err := transactionService.GetById(context.Background(), transactionDomain)

		assert.Error(t, err)
		assert.Equal(t, data, transactions.Domain{})

		transactionRepository.AssertExpectations(t)
	})

}

//
func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsCreate",
			mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()

		data, err := transactionService.Create(context.Background(), transactionDomain)

		assert.Nil(t, err)
		assert.Equal(t, data, transactionDomain)
		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Create - Error", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsCreate",
			mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()

		data, err := transactionService.Create(context.Background(), transactionDomain)

		assert.Error(t, err)
		assert.Equal(t, data, transactions.Domain{})
		transactionRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsUpdate",
			mock.Anything, mock.Anything).Return(transactionDomain, nil).Once()

		data, err := transactionService.Update(context.Background(), transactionDomain)

		assert.NotNil(t, data)
		assert.NoError(t, err)

		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update - Error (Update Fail)", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsUpdate",
			mock.Anything, mock.Anything).Return(transactions.Domain{}, businesses.ErrForTest).Once()

		data, err := transactionService.Update(context.Background(), transactionDomain)

		assert.Error(t, err)
		assert.Equal(t, data, transactions.Domain{})

		transactionRepository.AssertExpectations(t)
	})

}
func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsDelete",
			mock.Anything, mock.Anything).Return(nil).Once()
		err := transactionService.Delete(context.Background(), transactionDomain)

		assert.Nil(t, err)

		transactionRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Error", func(t *testing.T) {
		setup()
		transactionRepository.On("TransactionsDelete",
			mock.Anything, mock.Anything).Return(businesses.ErrForTest).Once()
		err := transactionService.Delete(context.Background(), transactionDomain)

		assert.Equal(t, err, businesses.ErrForTest)
		assert.Error(t, err)

		transactionRepository.AssertExpectations(t)
	})

}
