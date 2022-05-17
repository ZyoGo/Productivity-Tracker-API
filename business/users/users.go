package user

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	Id          string
	Username    string
	Password    string
	PhoneNumber int64
	CreatedAt   time.Time
	LastLogin   time.Time
	Deleted     bool
}

func NewUser(
	username string,
	password string,
	phoneNumber int64) Users {

	id := uuid.New()
	idString := strings.Replace(id.String(), "-", "", -1)

	return Users{
		Id:          idString,
		Username:    username,
		Password:    password,
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now(),
		LastLogin:   time.Now(),
		Deleted:     false,
	}
}

func (old *Users) ModifyUser(
	newUsername string,
	newPassword string,
	newPhoneNumber int64,
) Users {

	return Users{
		Id:          old.Id,
		Username:    newUsername,
		Password:    newPassword,
		PhoneNumber: newPhoneNumber,
		CreatedAt:   old.CreatedAt,
		LastLogin:   time.Now(),
		Deleted:     old.Deleted,
	}
}
