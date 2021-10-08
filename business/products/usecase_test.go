package products_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	businesses "mikro/business"
	"mikro/business/products"
	_mockProducts "mikro/business/products/mocks"
	_mockUsers "mikro/business/users/mocks"

	"testing"
	"time"
)

var productRepository _mockProducts.Repository
var userService _mockUsers.Usecase

var productService products.Usecase

var productDomain products.Domain
var listProductDomain []products.Domain

func setup() {
	productService = products.NewProductUsecase(&productRepository, &userService, time.Second*10)
	productDomain = products.Domain{
		ID:       1,
		SellerId: 1,
		Seller: products.Seller{
			StoreName: "Alfai Store",
			Name:      "A",
			City:      "Zimba",
		},
		Name:  "Brodi",
		Price: 1000000,
	}
	listProductDomain = append(listProductDomain, productDomain)

}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		productRepository.On("ProductsGetAll",
			mock.Anything).Return(listProductDomain, nil).Once()
		data, err := productService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listProductDomain))
		productRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2| UsersGetAll - Error", func(t *testing.T) {
		setup()
		productRepository.On("ProductsGetAll",
			mock.Anything).Return([]products.Domain{}, businesses.ErrForTest).Once()
		data, err := productService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []products.Domain{})
		productRepository.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | ProductsGetById - Success", func(t *testing.T) {
		setup()
		productRepository.On("ProductsGetById",
			mock.Anything, mock.Anything).Return(productDomain, nil).Once()
		data, err := productService.GetById(context.Background(), productDomain)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		productRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetById - Error", func(t *testing.T) {
		setup()
		productDomain.ID = 0
		productRepository.On("ProductsGetById",
			mock.Anything, mock.Anything).Return(products.Domain{}, businesses.ErrForTest).Once()
		data, err := productService.GetById(context.Background(), productDomain)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})

		productRepository.AssertExpectations(t)
	})

}

//
func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		productRepository.On("ProductsCreate",
			mock.Anything, mock.Anything).Return(productDomain, nil).Once()

		data, err := productService.Create(context.Background(), productDomain)

		assert.Nil(t, err)
		assert.Equal(t, data, productDomain)
		productRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Create - Error", func(t *testing.T) {
		setup()
		productRepository.On("ProductsCreate",
			mock.Anything, mock.Anything).Return(products.Domain{}, businesses.ErrForTest).Once()

		data, err := productService.Create(context.Background(), productDomain)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})
		productRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		productRepository.On("ProductsUpdate",
			mock.Anything, mock.Anything).Return(productDomain, nil).Once()

		data, err := productService.Update(context.Background(), productDomain)

		assert.NotNil(t, data)
		assert.NoError(t, err)

		productRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update - Error (Update Fail)", func(t *testing.T) {
		setup()
		productRepository.On("ProductsUpdate",
			mock.Anything, mock.Anything).Return(products.Domain{}, businesses.ErrForTest).Once()

		data, err := productService.Update(context.Background(), productDomain)

		assert.Error(t, err)
		assert.Equal(t, data, products.Domain{})

		productRepository.AssertExpectations(t)
	})

}
func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		productRepository.On("ProductsDelete",
			mock.Anything, mock.Anything).Return(nil).Once()
		err := productService.Delete(context.Background(), productDomain)

		assert.Nil(t, err)

		productRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Error", func(t *testing.T) {
		setup()
		productRepository.On("ProductsDelete",
			mock.Anything, mock.Anything).Return(businesses.ErrForTest).Once()
		err := productService.Delete(context.Background(), productDomain)

		assert.Equal(t, err, businesses.ErrForTest)
		assert.Error(t, err)

		productRepository.AssertExpectations(t)
	})

}
