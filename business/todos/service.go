package todos

import (
	"github.com/go-playground/validator/v10"
	"github.com/w33h/Productivity-Tracker-API/business/todos/spec"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	"github.com/w33h/Productivity-Tracker-API/exception"
)

type RepositoryTodos interface {
	InsertTodo(todo Todo) (id string, err error)
	UpdateTodo(todo Todo) (err error)
	DeleteTodo(id string) (err error)
	FindByStatus(status string) (todo []Todo, err error)
	FindById(id string) (todo *Todo, err error)
	FindAllTodo() (todo []Todo, err error)
}

type ServiceTodos interface {
	CreateTodo(specTodo spec.UpsertTodosSpec) (id string, err error)
	UpdateTodo(specTodo spec.UpsertTodosSpec, id string) (err error)
	DeleteTodo(id string) (err error)
	GetByStatus(status string) (todo []Todo, err error)
	GetById(id string) (todo *Todo, err error)
	GetAllTodo() (todo []Todo, err error)
}

type todoService struct {
	todoRepo RepositoryTodos
	userRepo user.RepositoryUser
	validate *validator.Validate
}

func NewTodoService(repoTodo RepositoryTodos, repoUser user.RepositoryUser) ServiceTodos {
	return &todoService{
		todoRepo: repoTodo,
		userRepo: repoUser,
		validate: validator.New(),
	}
}

func (s *todoService) CreateTodo(specTodo spec.UpsertTodosSpec) (id string, err error) {
	err = s.validate.Struct(specTodo)
	if err != nil {
		return id, exception.ErrInvalidSpec
	}

	_, err = s.userRepo.FindById(specTodo.UserId)
	if err != nil {
		return id, exception.ErrNotFound
	}

	newTodo := NewTodos(specTodo.Status, specTodo.Content, specTodo.UserId)

	id, err = s.todoRepo.InsertTodo(newTodo)
	if err != nil {
		return id, exception.ErrInternalServer
	}

	return id, nil
}

func (s *todoService) UpdateTodo(specTodo spec.UpsertTodosSpec, id string) (err error) {
	err = s.validate.Struct(specTodo)
	if err != nil {
		return exception.ErrInvalidSpec
	}

	oldTodo, err := s.todoRepo.FindById(id)
	if err != nil {
		return exception.ErrNotFound
	}

	newTodo := oldTodo.ModifyTodos(specTodo.Content, specTodo.Status)

	err = s.todoRepo.UpdateTodo(newTodo)
	if err != nil {
		return exception.ErrInternalServer
	}

	return nil
}

func (s *todoService) DeleteTodo(id string) (err error) {
	err = s.todoRepo.DeleteTodo(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *todoService) GetByStatus(status string) (todo []Todo, err error) {
	todo, err = s.todoRepo.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *todoService) GetById(id string) (todo *Todo, err error) {
	todo, err = s.todoRepo.FindById(id)
	if err != nil {
		return nil, exception.ErrNotFound
	}

	return todo, nil
}

func (s *todoService) GetAllTodo() (todo []Todo, err error) {
	todo, err = s.todoRepo.FindAllTodo()
	if err != nil {
		return nil, exception.ErrInternalServer
	}

	return todo, nil
}
