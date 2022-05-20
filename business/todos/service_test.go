package todos_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.com/w33h/Productivity-Tracker-API/business/todos"
	"github.com/w33h/Productivity-Tracker-API/business/todos/spec"
	"github.com/w33h/Productivity-Tracker-API/exception"
	todoMock "github.com/w33h/Productivity-Tracker-API/repository/todos/mocks"
	userMock "github.com/w33h/Productivity-Tracker-API/repository/users/mocks"
	"testing"
	"time"
)

type UnitTestSuite struct {
	suite.Suite
	todoService todos.ServiceTodos
	todoMock    *todoMock.RepositoryTodos
	userMock    *userMock.RepositoryUser
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (s *UnitTestSuite) SetupTest() {
	todosMock := todoMock.RepositoryTodos{}
	usersMock := userMock.RepositoryUser{}
	todosService := todos.NewTodoService(&todosMock, &usersMock)

	s.todoMock = &todosMock
	s.todoService = todosService
	s.userMock = &usersMock
}

func (s *UnitTestSuite) TestTodoService_GetById_Success() {
	responseMock := todos.Todo{
		Id:        uuid.New().String(),
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	s.todoMock.On("FindById", mock.Anything).Return(&responseMock, nil)
	result, err := s.todoService.GetById("5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e", responseMock.UserId)
	fmt.Println("error =", err)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_GetById_Failed_InvalidID() {
	s.todoMock.On("FindById", "1").Return(nil, exception.ErrNotFound)
	result, err := s.todoService.GetById("1", "1")

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_GetByStatus_Success() {
	responseMock := []todos.Todo{
		{
			Id:        "1",
			UserId:    "ba7ae1fd-ce80-450c-b68b-8a2502ee5283",
			Status:    "InProgress",
			Content:   "A",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Deleted:   false,
		},
	}
	s.todoMock.On("FindByStatus", "InProgress").Return(responseMock, nil)
	result, err := s.todoService.GetByStatus("InProgress")
	fmt.Println("result = ", result)
	fmt.Println("error = ", err)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_GetByStatus_Failed_NotFoundStatus() {
	s.todoMock.On("FindByStatus", "InProgress").Return(nil, exception.ErrNotFound)
	result, err := s.todoService.GetByStatus("InProgress")

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_GetAllTodo_Success() {
	responseMock := []todos.Todo{
		{
			Id:        "1",
			UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
			Status:    "InProgress",
			Content:   "A",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Deleted:   false,
		},
		{
			Id:        "2",
			UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
			Status:    "Completed",
			Content:   "B",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Deleted:   false,
		},
	}
	s.todoMock.On("FindAllTodo", mock.Anything).Return(responseMock, nil)
	result, err := s.todoService.GetAllTodo("5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e")

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_GetAllTodo_Failed_NotFoundTodo() {
	s.todoMock.On("FindAllTodo", mock.Anything).Return(nil, exception.ErrNotFound)
	result, err := s.todoService.GetAllTodo("1")

	assert.Error(s.T(), err)
	assert.Nil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_DeleteTodo_Success() {
	s.todoMock.On("DeleteTodo", mock.Anything).Return(nil)
	err := s.todoService.DeleteTodo("1")

	assert.NoError(s.T(), err)
}

func (s *UnitTestSuite) TestTodoService_DeleteTodo_Failed_InvalidID() {
	s.todoMock.On("DeleteTodo", mock.Anything).Return(exception.ErrNotFound)
	err := s.todoService.DeleteTodo("1")

	assert.Error(s.T(), err)
}

func (s *UnitTestSuite) TestTodoService_CreateTodo_Success() {
	responseMock := todos.Todo{
		Id:        uuid.New().String(),
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	specTodo := spec.UpsertTodosSpec{
		UserId:  "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Content: "Todo A",
		Status:  "InProgress",
	}

	s.userMock.On("FindById", mock.Anything).Return(nil, nil)
	s.todoMock.On("InsertTodo", mock.Anything).Return(responseMock.Id, nil)

	result, err := s.todoService.CreateTodo(specTodo, responseMock.UserId)

	assert.NoError(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_CreateTodo_Failed_InvalidSpec() {
	responseMock := todos.Todo{
		Id:        uuid.New().String(),
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	specTodo := spec.UpsertTodosSpec{
		UserId:  "5e0efcb7-9176-4526-a5a4-b82a01f4b1e",
		Content: "Todo A",
		Status:  "InProgress",
	}

	s.userMock.On("FindById", mock.Anything).Return(nil, nil)
	s.todoMock.On("InsertTodo", mock.Anything).Return(responseMock.Id, nil)

	result, err := s.todoService.CreateTodo(specTodo, responseMock.UserId)

	assert.Error(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_CreateTodo_Failed_IdNotFound() {
	responseMock := todos.Todo{
		Id:        uuid.New().String(),
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	specTodo := spec.UpsertTodosSpec{
		UserId:  "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Content: "Todo A",
		Status:  "InProgress",
	}

	s.userMock.On("FindById", mock.Anything).Return(nil, exception.ErrNotFound)
	s.todoMock.On("InsertTodo", mock.Anything).Return(responseMock.Id, nil)

	result, err := s.todoService.CreateTodo(specTodo, responseMock.UserId)

	assert.Error(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_CreateTodo_Failed_InternalServerError() {
	responseMock := todos.Todo{
		Id:        uuid.New().String(),
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	specTodo := spec.UpsertTodosSpec{
		UserId:  "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Content: "Todo A",
		Status:  "InProgress",
	}

	s.userMock.On("FindById", mock.Anything).Return(nil, nil)
	s.todoMock.On("InsertTodo", mock.Anything).Return(responseMock.Id, exception.ErrInternalServer)

	result, err := s.todoService.CreateTodo(specTodo, responseMock.UserId)

	assert.Error(s.T(), err)
	assert.NotNil(s.T(), result)
}

func (s *UnitTestSuite) TestTodoService_UpdateTodo_Success() {
	responseMock := todos.Todo{
		Id:        "1",
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	specTodo := spec.UpsertTodosSpec{
		UserId:  "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Content: "Todo A",
		Status:  "InProgress",
	}

	s.todoMock.On("FindById", mock.Anything).Return(&responseMock, nil)
	s.todoMock.On("UpdateTodo", mock.Anything).Return(nil)

	err := s.todoService.UpdateTodo(specTodo, responseMock.Id)

	assert.NoError(s.T(), err)
}

func (s *UnitTestSuite) TestTodoService_UpdateTodo_Failed_InvalidID() {
	responseMock := todos.Todo{
		Id:        "1",
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	specTodo := spec.UpsertTodosSpec{
		UserId:  "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Content: "Todo A",
		Status:  "InProgress",
	}

	s.todoMock.On("FindById", mock.Anything).Return(nil, exception.ErrNotFound)
	s.todoMock.On("UpdateTodo", mock.Anything).Return(nil)

	err := s.todoService.UpdateTodo(specTodo, responseMock.Id)

	assert.Error(s.T(), err)
}

func (s *UnitTestSuite) TestTodoService_UpdateTodo_Failed_InvalidSpec() {
	responseMock := todos.Todo{
		Id:        "1",
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	specTodo := spec.UpsertTodosSpec{
		UserId:  "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1",
		Content: "Todo A",
		Status:  "InProgress",
	}

	s.todoMock.On("FindById", mock.Anything).Return(&responseMock, nil)
	s.todoMock.On("UpdateTodo", mock.Anything).Return(exception.ErrInternalServer)

	err := s.todoService.UpdateTodo(specTodo, responseMock.Id)

	assert.Error(s.T(), err)
}

func (s *UnitTestSuite) TestTodoService_UpdateTodo_Failed_InternalServerError() {
	responseMock := todos.Todo{
		Id:        "1",
		UserId:    "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Status:    "InProgress",
		Content:   "A",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Deleted:   false,
	}
	specTodo := spec.UpsertTodosSpec{
		UserId:  "5e0efcb7-9176-4526-a5a4-b82a0e1f4b1e",
		Content: "Todo A",
		Status:  "InProgress",
	}

	s.todoMock.On("FindById", mock.Anything).Return(&responseMock, nil)
	s.todoMock.On("UpdateTodo", mock.Anything).Return(exception.ErrInternalServer)

	err := s.todoService.UpdateTodo(specTodo, responseMock.Id)

	assert.Error(s.T(), err)
}
