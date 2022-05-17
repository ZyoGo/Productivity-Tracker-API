package user

import (
	domain "github.com/w33h/Productivity-Tracker-API/business/users"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.RepositoryUser {
	return &userRepository{
		db,
	}
}

func (r *userRepository) FindById(id string) (result *domain.Users, err error) {
	if err = r.db.Where("Id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	result.Id = id

	return result, nil
}

func (r *userRepository) InsertUser(user domain.Users) (id domain.Users, err error) {
	err = r.db.Create(&user).Error

	if err != nil {
		return id, err
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user *domain.Users) (err error) {
	if err = r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) DeleteUser(id string) (err error) {
	err = r.db.Where("id = ?", id).First(domain.Users{}).Delete(domain.Users{}).Error

	if err != nil {
		return err
	}

	return nil
}
