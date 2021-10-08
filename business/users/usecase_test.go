package users_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mikro/app/middlewares"
	businesses "mikro/business"
	"mikro/business/users"
	_mockUser "mikro/business/users/mocks"
	"testing"
	"time"
)

var userRepository _mockUser.Repository
var userService users.Usecase

var userDomain users.Domain
var listUserDomain []users.Domain
var token string

func setup() {
	userService = users.NewUserUsecase(&userRepository, time.Second*10, &middlewares.ConfigJWT{})
	userDomain = users.Domain{
		ID:          1,
		Name:        "dsad",
		StoreName:   "dsa",
		City:        "sad",
		Phone:       "dsadas",
		Username:    "dsdsad",
		IsAdmin:     false,
		BankAccount: "dsa",
		NoAccount:   "dsadsa",
		Password:    "$2a$12$SaZVHXYiMiygeHoVV33cY..FwaM/oFNO9EVTnscfnKOlKvIWtnRRS",
		Token:       "sadasdsa",
	}
	token = "token"
	listUserDomain = append(listUserDomain, userDomain)

}

func TestGetAll(t *testing.T) {
	t.Run("Test Case 1 | GetAll - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetAll",
			mock.Anything).Return(listUserDomain, nil).Once()
		data, err := userService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.Equal(t, len(data), len(listUserDomain))
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2| UsersGetAll - Error", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetAll",
			mock.Anything).Return([]users.Domain{}, businesses.ErrForTest).Once()
		data, err := userService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, data, []users.Domain{})
		userRepository.AssertExpectations(t)
	})
}

func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | GetById - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetById",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		data, err := userService.GetById(context.Background(), userDomain)

		assert.NoError(t, err)
		assert.NotNil(t, data)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | GetById - Error", func(t *testing.T) {
		setup()
		userDomain.ID = 0
		userRepository.On("UsersGetById",
			mock.Anything, mock.Anything).Return(users.Domain{}, businesses.ErrForTest).Once()
		data, err := userService.GetById(context.Background(), userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})

		userRepository.AssertExpectations(t)
	})

}

func TestCreate(t *testing.T) {
	t.Run("Test Case 1 | Create - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByUsername",
			mock.Anything, mock.Anything).Return(users.Domain{}, nil).Once()
		userRepository.On("UsersCreate",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()

		data, err := userService.Create(context.Background(), userDomain)

		assert.Nil(t, err)
		assert.Equal(t, data, userDomain)
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Create - Error", func(t *testing.T) {
		setup()

		_, err := userService.Create(context.Background(), users.Domain{})

		assert.Error(t, err)
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Create - Error", func(t *testing.T) {
		setup()

		_, err := userService.Create(context.Background(), users.Domain{
			Username: "er",
		})

		assert.Error(t, err)
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Error - Email Has Been Used", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByUsername",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		data, err := userService.Create(context.Background(), userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 3 | Error - Date Not Valid", func(t *testing.T) {
		userRepository.On("UsersGetByUsername",
			mock.Anything, mock.Anything).Return(users.Domain{}, nil).Once()
		userRepository.On("UsersCreate",
			mock.Anything, mock.Anything).Return(users.Domain{}, businesses.ErrForTest).Once()

		data, err := userService.Create(context.Background(), userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})
		userRepository.AssertExpectations(t)
	})
}

//
func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Login - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByUsername",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		_, err := userService.Login(context.Background(), users.Domain{
			Username: "dsdsad",
			Password: "123",
		})

		assert.NotNil(t, token)
		assert.NoError(t, err)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Login - Error (Username/Pass Not Found)", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByUsername",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		data, err := userService.Login(context.Background(), userDomain)

		assert.Equal(t, users.Domain{}, data)
		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test Case 1 | Update - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByUsername",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		userRepository.On("UsersUpdate",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		data, err := userService.Update(context.Background(), userDomain)

		assert.NotNil(t, data)
		assert.NoError(t, err)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Update - Error (Update Fail)", func(t *testing.T) {
		setup()
		userRepository.On("UsersGetByUsername",
			mock.Anything, mock.Anything).Return(userDomain, nil).Once()
		userRepository.On("UsersUpdate",
			mock.Anything, mock.Anything).Return(users.Domain{}, businesses.ErrForTest).Once()
		data, err := userService.Update(context.Background(), userDomain)

		assert.Error(t, err)
		assert.Equal(t, data, users.Domain{})

		userRepository.AssertExpectations(t)
	})

}
func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Delete - Success", func(t *testing.T) {
		setup()
		userRepository.On("UsersDelete",
			mock.Anything, mock.Anything).Return(nil).Once()
		err := userService.Delete(context.Background(), userDomain)

		assert.Nil(t, err)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test Case 2 | Delete - Error", func(t *testing.T) {
		setup()
		userRepository.On("UsersDelete",
			mock.Anything, mock.Anything).Return(businesses.ErrForTest).Once()
		err := userService.Delete(context.Background(), userDomain)

		assert.Equal(t, err, businesses.ErrForTest)
		assert.Error(t, err)

		userRepository.AssertExpectations(t)
	})

}
