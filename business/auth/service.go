package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/w33h/Productivity-Tracker-API/business/auth/spec"
	"github.com/w33h/Productivity-Tracker-API/exception"
)

type RepositoryAuth interface {
	VerifyCredential(username, password string) (userId string, err error)
}

type ServiceAuth interface {
	LoginUser(loginSpec spec.UpsertAuthSpec) (auth *Auth, err error)
}

type serviceAuth struct {
	repo     RepositoryAuth
	validate *validator.Validate
}

func NewAuthService(repo RepositoryAuth) ServiceAuth {
	return &serviceAuth{
		repo:     repo,
		validate: validator.New(),
	}
}

func (s *serviceAuth) LoginUser(loginSpec spec.UpsertAuthSpec) (auth *Auth, err error) {
	err = s.validate.Struct(loginSpec)
	if err != nil {
		return nil, exception.ErrInvalidSpec
	}

	id, err := s.repo.VerifyCredential(loginSpec.Username, loginSpec.Password)
	if err != nil {
		return nil, exception.ErrNotFound
	}
	fmt.Println("id = ", id)

	expirationTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		Username: loginSpec.Username,
		UserId:   id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("Secret_JWT")

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return nil, exception.ErrInternalServer
	}

	auth = &Auth{
		Token:  tokenString,
		UserID: id,
	}

	return auth, nil
}
