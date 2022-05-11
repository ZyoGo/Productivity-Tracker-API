package user

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/w33h/Productivity-Tracker-API/business/users/spec"
)

type RepositoryUser interface {
	Get(id int32) (result *Users, err error)
	Create(user Users) (id Users, err error)
	Update(user Users) (result Users, err error)
	Delete(id int32) (err error)
}

type ServiceUser interface {
	Get(id int32) (user *Users, err error)
	Create(userSpec spec.UpsertUserSpec) (id Users, err error)
	Update(id int32, userSpec spec.UpsertUserSpec) (result Users, err error)
	Delete(id int32) (err error)
}

type userService struct {
	repo     RepositoryUser
	validate *validator.Validate
}

func NewUserService(repo RepositoryUser) ServiceUser {
	return &userService{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *userService) Get(id int32) (user *Users, err error) {
	user, err = s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Create(userSpec spec.UpsertUserSpec) (id Users, err error) {
	err = s.validate.Struct(&userSpec)

	if err != nil {
		return id, errors.New("validate")
	}

	id = NewUser(userSpec.Username, userSpec.Password, userSpec.PhoneNumber)

	id, err = s.repo.Create(id)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (s *userService) Update(id int32, userSpec spec.UpsertUserSpec) (result Users, err error) {
	err = s.validate.Struct(&userSpec)

	if err != nil {
		return result, errors.New("validate error")
	}

	oldUser, err := s.repo.Get(id)

	if err != nil {
		return result, errors.New("user not found")
	}

	result = oldUser.ModifyUser(userSpec.Username, userSpec.Password, userSpec.PhoneNumber)

	result, err = s.repo.Update(result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *userService) Delete(id int32) (err error) {
	err = s.repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

