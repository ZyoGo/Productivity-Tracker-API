package auth

import (
	"github.com/w33h/Productivity-Tracker-API/business/auth"
	domain "github.com/w33h/Productivity-Tracker-API/business/users"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.RepositoryAuth {
	return &userRepository{db}
}

func (r *userRepository) VerifyCredential(username, password string) (id string, err error) {
	var credential domain.Users
	err = r.db.Table("users").Where("username = ? AND password = ?", username, password).First(&credential).Error

	if err != nil {
		return id, err
	}

	id = credential.Id

	return id, nil
}
