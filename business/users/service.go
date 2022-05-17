package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/w33h/Productivity-Tracker-API/business/users/spec"
)

type RepositoryUser interface {
	FindById(id string) (result *Users, err error)
	InsertUser(user Users) (id Users, err error)
	UpdateUser(user *Users) (err error)
	DeleteUser(id string) (err error)
}

type ServiceUser interface {
	GetById(id string) (user *Users, err error)
	CreateUser(userSpec spec.UpsertUserSpec) (user Users, err error)
	UpdateUser(userSpec spec.UpsertUserSpec, id string) (err error)
	DeleteUser(id string) (err error)
}

type userService struct {
	userRepo RepositoryUser
	validate *validator.Validate
}

func NewUserService(repo RepositoryUser) ServiceUser {
	return &userService{
		userRepo: repo,
		validate: validator.New(),
	}
}

func (s *userService) GetById(id string) (user *Users, err error) {
	user, err = s.userRepo.FindById(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(userSpec spec.UpsertUserSpec) (user Users, err error) {
	err = s.validate.Struct(&userSpec)

	if err != nil {
		return user, err
	}

	newUser := NewUser(userSpec.Username, userSpec.Password, userSpec.PhoneNumber)

	user, err = s.userRepo.InsertUser(newUser)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) UpdateUser(userSpec spec.UpsertUserSpec, id string) (err error) {
	err = s.validate.Struct(&userSpec)
	if err != nil {
		return err
	}

	user, err := s.userRepo.FindById(id)
	if err != nil {
		return err
	}

	newUser := user.ModifyUser(userSpec.Username, userSpec.Password, userSpec.PhoneNumber)

	err = s.userRepo.UpdateUser(&newUser)
	if err != nil {
		return err
	}

	return nil
}

func (s *userService) DeleteUser(id string) (err error) {
	err = s.userRepo.DeleteUser(id)
	fmt.Println("id delete = ", id)
	fmt.Println("err id delete = ", err)

	if err != nil {
		return err
	}

	return nil
}
