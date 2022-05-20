package request

import "github.com/w33h/Productivity-Tracker-API/business/auth/spec"

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (req *AuthRequest) ToSpec() *spec.UpsertAuthSpec {
	return &spec.UpsertAuthSpec{
		Username: req.Username,
		Password: req.Password,
	}
}