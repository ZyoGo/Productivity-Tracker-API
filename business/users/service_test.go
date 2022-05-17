package user_test

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/w33h/Productivity-Tracker-API/business/users/spec"
	"github.com/w33h/Productivity-Tracker-API/repository/users/mocks"
	"testing"
	"time"

	user "github.com/w33h/Productivity-Tracker-API/business/users"
)

type UnitTestSuite struct {
	suite.Suite
	userService user.ServiceUser
	userMock *mocks.RepositoryUser
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (s *UnitTestSuite) SetupTest()  {
	userMocks := mocks.RepositoryUser{}
	userServices := user.NewUserService(&userMocks)

	s.userService = userServices
	s.userMock = &userMocks
}

func (s *UnitTestSuite) TestUserService_Delete_Success()  {
		s.userMock.On("DeleteUser", mock.Anything).Return(nil)
		err := s.userService.DeleteUser("1")

		assert.NoError(s.T(), err)
}

func (s *UnitTestSuite) TestUserService_Delete_Failed()  {
	s.userMock.On("DeleteUser", mock.Anything).Return(errors.New("Failed"))
	err := s.userService.DeleteUser("1")

	assert.Error(s.T(), err)
}

func (s *UnitTestSuite) TestUserService_Create_Success()  {
	validRequest := user.Users{
					Id:          uuid.New().String(),
					Username:    "userA@gmail.com",
					Password:    "userA123",
					PhoneNumber: +6280000000000,
					CreatedAt:   time.Now(),
					LastLogin:   time.Now(),
					Deleted:     false,
				}
	userSpec := spec.UpsertUserSpec{
					Username:    "userA@gmail.com",
					Password:    "userA123",
					PhoneNumber: +6280000000000,
				}

	s.userMock.On("InsertUser", mock.Anything).Return(validRequest, nil)
	result, err := s.userService.CreateUser(userSpec)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestUserService_Create_InternalServerFailed()  {
	invalidRequest := user.Users{}
	userSpec := spec.UpsertUserSpec{
		Username:    "userA@gmail.com",
		Password:    "userA123",
		PhoneNumber: +6280000000000,
	}

	s.userMock.On("InsertUser", mock.Anything).Return(invalidRequest, errors.New("internal server failed"))
	result, err := s.userService.CreateUser(userSpec)

	assert.Error(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestUserService_Create_BadRequest()  {
	invalidRequest := user.Users{
		Id:          "1",
		Username:    "userA@gmail.com",
		Password:    "userA123",
		PhoneNumber: +6280000000000,
		CreatedAt:   time.Now(),
		LastLogin:   time.Now(),
		Deleted:     false,
	}
	userSpec := spec.UpsertUserSpec{
		Username:    "userA",
		Password:    "userA123",
		PhoneNumber: +6280000000000,
	}

	s.userMock.On("InsertUser", mock.Anything).Return(invalidRequest, errors.New("user spec failed"))
	result, err := s.userService.CreateUser(userSpec)

	assert.Error(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestUserService_Get_Success() {
	s.userMock.On("FindById", "1").Return(&user.Users{}, nil)
	result, err := s.userService.GetById("1")

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestUserService_Get_Failed_InvalidID() {
	s.userMock.On("FindById", "1").Return(nil, errors.New("id not found"))
	result, err := s.userService.GetById("1")

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
}

func (s *UnitTestSuite) TestUserService_Update_Success() {
	invalidRequest := user.Users{
		Id:          "1",
		Username:    "userA@gmail.com",
		Password:    "userA123",
		PhoneNumber: +6280000000000,
		CreatedAt:   time.Now(),
		LastLogin:   time.Now(),
		Deleted:     false,
	}
	userSpec := spec.UpsertUserSpec{
		Username:    "userA@gmail.com",
		Password:    "userA123",
		PhoneNumber: +6280000000000,
	}

	s.userMock.On("FindById", mock.Anything).Return(&invalidRequest, nil)
	s.userMock.On("UpdateUser", mock.Anything).Return(nil)

	err := s.userService.UpdateUser(userSpec, "1")

	assert.NoError(s.T(), err)
}

func (s *UnitTestSuite) TestUserService_Update_Failed_InternalServerFailed() {
	userSpec := spec.UpsertUserSpec{
		Username:    "userA@gmail.com",
		Password:    "userA123",
		PhoneNumber: +6280000000000,
	}

	s.userMock.On("FindById", mock.Anything).Return(nil, errors.New("id not found"))
	s.userMock.On("UpdateUser", mock.Anything).Return(nil)

	err := s.userService.UpdateUser(userSpec, "1")

	assert.Error(s.T(), err)
}

func (s *UnitTestSuite) TestUserService_Update_Failed_BadRequest() {
	invalidRequest := user.Users{
		Id:          "1",
		Username:    "userA@gmail.com",
		Password:    "userA123",
		PhoneNumber: +6280000000000,
		CreatedAt:   time.Now(),
		LastLogin:   time.Now(),
		Deleted:     false,
	}
	userSpec := spec.UpsertUserSpec{
		Username:    "userA",
		Password:    "userA123",
		PhoneNumber: +6280000000000,
	}

	s.userMock.On("FindById", mock.Anything).Return(&invalidRequest, nil)
	s.userMock.On("UpdateUser", mock.Anything).Return(errors.New("bad request"))

	err := s.userService.UpdateUser(userSpec, "1")

	assert.Error(s.T(), err)
}
