package mock

import (
	"github.com/stretchr/testify/mock"
	domain "github.com/w33h/Productivity-Tracker-API/business/users"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (m *UserRepositoryMock) Get(id int32) (result *domain.Users, err error) {
	args := m.Mock.Called(id)

	if args.Get(0) == nil {
		return result, err
	}

	result = args.Get(0).(*domain.Users)
	return result, nil
}

func (m *UserRepositoryMock) Create(user domain.Users) (id domain.Users, err error) {
	args := m.Mock.Called(user)

	if args.Get(0) == nil {
		return id, err
	}

	id = args.Get(0).(domain.Users)
	return id, nil
}

func (m *UserRepositoryMock) Update(user domain.Users) (result domain.Users, err error) {
	args := m.Mock.Called(user)

	if args.Get(0) == nil {
		return result, err
	}

	return result, nil
}

func (m *UserRepositoryMock) Delete(id int32) (err error) {
	args := m.Mock.Called(id)

	if args.Get(0) == nil {
		return err
	}

	return nil
}
