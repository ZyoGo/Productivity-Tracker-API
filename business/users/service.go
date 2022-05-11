package user

import (
	"github.com/go-playground/validator/v10"
)

type UserRepository interface {
	Get(id int32) (user *Users, err error)
	Create(user Users) (id Users, err error)
	Update(user Users) (err error)
	Delete(id int32) (err error)
}

type UserService interface {
	Get(id int32) (user *Users, err error)
	Create(user Users) (id Users, err error)
	Update(user Users) (err error)
	Delete(user Users) (err error)
}

type userService struct {
	userRepo UserRepository
	validate *validator.Validate
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		userRepo: repo,
		validate: validator.New(),
	}
}
