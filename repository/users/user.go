package user

import (
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.RepositoryUser {
	return &userRepository{
		db,
	}
}

func (r *userRepository) Get(id int32) (result *user.Users, err error) {
	err = r.db.Where("Id = ?", id).Error

	if err != nil {
		return nil, err
	}

	result.Id = int64(id)

	return result, nil
}

func (r *userRepository) Create(user user.Users) (id user.Users, err error) {
	err = r.db.Create(&user).Error

	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *userRepository) Update(user user.Users) (result user.Users, err error) {
	err = r.db.Save(&user).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *userRepository) Delete(id int32) (err error) {
	err = r.db.Delete(id).Error

	if err != nil {
		return err
	}

	return nil
}
