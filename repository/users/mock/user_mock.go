package mock

import (
	"github.com/stretchr/testify/mock"
	domain "github.com/w33h/Productivity-Tracker-API/business/users"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (m *UserRepositoryMock) FindById(id string) (result *domain.Users, err error) {
	args := m.Mock.Called(id)

	if args.Get(0) == nil {
		return nil, err
	}

	// result = args.Get(0).(*domain.Users)
	user := args.Get(0).(domain.Users)
	result = &user
	return result, nil
}

func (m *UserRepositoryMock) InsertUser(user domain.Users) (id domain.Users, err error) {
	args := m.Mock.Called(user)

	if args.Get(0) == nil {
		return id, err
	}

	id = args.Get(0).(domain.Users)
	return id, nil
}

func (m *UserRepositoryMock) UpdateUser(user *domain.Users) (err error) {
	args := m.Mock.Called(user)

	if args.Get(0) == nil {
		return err
	}

	return nil
}

func (m *UserRepositoryMock) DeleteUser(id string) (err error) {
	args := m.Mock.Called(id)

	if args.Get(0) == nil {
		return err
	}

	return nil
}
