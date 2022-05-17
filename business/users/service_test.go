package user_test

import (
	"errors"
	"github.com/w33h/Productivity-Tracker-API/repository/users/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	"github.com/w33h/Productivity-Tracker-API/business/users/spec"
)

var userRepository = mocks.RepositoryUser{}
var userService = user.NewUserService(&userRepository)

func TestUserService_UpdateUser(t *testing.T) {
	t.Run("Update user with valid request", func(t *testing.T) {
		validRequest := user.Users{
			Id:          "1",
			Username:    "user1@gmail.com",
			Password:    "user123456",
			PhoneNumber: 628222222222,
			CreatedAt:   time.Now(),
			LastLogin:   time.Now(),
			Deleted:     false,
		}

		userSpec := spec.UpsertUserSpec{
			Username:    "user123@gmail.com",
			Password:    "user123456",
			PhoneNumber: 6282394707112,
		}

		userRepository.On("FindById", mock.Anything).Return(&validRequest, nil)
		userRepository.On("UpdateUser", mock.Anything).Return(nil)
		err := userService.UpdateUser(userSpec, "1")
		//fmt.Println("error = ", err.Error)
		//expectedError := errors.New("test")

		assert.NoError(t, err)
		//assert.Equal(t, expectedError, err)
	})
}

func TestUserService_Delete(t *testing.T) {
	t.Run("Delete user with invalid id", func(t *testing.T) {
		userRepository.On("DeleteUser", mock.Anything).Return(errors.New("failed")).Once()
		err := userService.DeleteUser("1")

		assert.Error(t, err)
	})

	t.Run("Delete user with valid id", func(t *testing.T) {
		userRepository.On("DeleteUser", mock.Anything).Return(nil)
		err := userService.DeleteUser("1")

		assert.NoError(t, err)
	})
}

func TestUserService_Create(t *testing.T) {
	t.Run("Create user with valid request", func(t *testing.T) {
		validRequest := user.Users{
			Id:          uuid.New().String(),
			Username:    "usera@gmail.com",
			Password:    "usera123",
			PhoneNumber: +6282394707112,
			CreatedAt:   time.Now(),
			LastLogin:   time.Now(),
			Deleted:     false,
		}

		userSpec := spec.UpsertUserSpec{
			Username:    "usera@gmail.com",
			Password:    "usera123",
			PhoneNumber: +6282394707112,
		}

		userRepository.Mock.On("InsertUser", mock.Anything).Return(validRequest, nil)
		result, err := userService.CreateUser(userSpec)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Create user with invalid request", func(t *testing.T) {
		invalidResponse := user.Users{}

		userSpecInvalid := spec.UpsertUserSpec{
			Username:    "usera@gmail.com",
			Password:    "123",
			PhoneNumber: +6282394707112,
		}

		userRepository.Mock.On("InsertUser", mock.Anything).Return(invalidResponse, nil)
		result, err := userService.CreateUser(userSpecInvalid)

		assert.Error(t, err)
		assert.NotNil(t, result)
	})
}

func TestUserService_Get(t *testing.T) {
	t.Run("Get user with valid id", func(t *testing.T) {
		validRequest := user.Users{
			Id:          "1",
			Username:    "user1@gmail.com",
			Password:    "user123456",
			PhoneNumber: +628222222222,
			CreatedAt:   time.Now(),
			LastLogin:   time.Now(),
			Deleted:     false,
		}

		userRepository.On("FindById", mock.Anything).Return(&validRequest, nil)
		result, err := userService.GetById("1")

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Get user with invalid id", func(t *testing.T) {
		userRepository.On("FindById", "12").Return(nil, errors.New("failed"))
		resultService, err := userService.GetById("12")

		assert.Error(t, err)
		assert.Nil(t, resultService)
	})
}
