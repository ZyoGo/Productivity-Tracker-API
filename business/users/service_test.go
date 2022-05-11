package user_test

import (
	"encoding/binary"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	"github.com/w33h/Productivity-Tracker-API/business/users/spec"
	"github.com/w33h/Productivity-Tracker-API/repository/users/mock"
	"testing"
	"time"
)

var userRepository = &mock.UserRepositoryMock{}
var userService = user.NewUserService(userRepository)

func TestUserService_Create(t *testing.T) {
	t.Run("Create user with valid request", func(t *testing.T) {
		validRequest := user.Users{
			Id:           uuid.New(),
			Username:     "usera@gmail.com",
			Password:     "usera123",
			PhoneNumber: +6282394707112,
			CreatedAt:   time.Time{},
			LastLogin:   time.Time{},
			Deleted:      false,
		}

		userSpec := spec.UpsertUserSpec{
			Username:     "usera@gmail.com",
			Password:     "usera123",
			PhoneNumber: +6282394707112,
		}

		userRepository.Mock.On("Create", mock2.Anything).Return(validRequest, nil)
		result, err := userService.Create(userSpec)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("Create user with invalid request", func(t *testing.T) {
		invalidResponse := user.Users{}

		userSpecInvalid := spec.UpsertUserSpec{
			Username:     "usera@gmail.com",
			Password:     "123",
			PhoneNumber: +6282394707112,
		}

		userRepository.Mock.On("Create", mock2.Anything).Return(invalidResponse, nil)
		result, err := userService.Create(userSpecInvalid)

		assert.Error(t, err)
		assert.NotNil(t, result)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("Get user with valid id", func(t *testing.T) {
		validRequest := user.Users{
			Id:           uuid,
			Username:     "usera@gmail.com",
			Password:     "usera123",
			PhoneNumber: +6282394707112,
			CreatedAt:   time.Time{},
			LastLogin:   time.Time{},
			Deleted:      false,
		}

		userRepository.Mock.On("Get", 12).Return(validRequest, nil)
	})
}