package request

import "github.com/w33h/Productivity-Tracker-API/business/users/spec"

type CreateRequestUser struct {
	Username    string `json:"username" validate:"required,email"`
	Password    string `json:"password" validate:"min=5"`
	PhoneNumber int64  `json:"phone_number" validate:"number"`
}

func (req *CreateRequestUser) ToSpecUser() *spec.UpsertUserSpec {
	return &spec.UpsertUserSpec{
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
	}
}
