package user

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

type Users struct {
	Id          int64
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
	idString := id.String()
	newId := strings.Replace(idString, "-", "", 1)

	return Users{
		Id:          strconv.,
		Username:    username,
		Password:    password,
		PhoneNumber: phoneNumber,
		CreatedAt:   time.Now(),
		LastLogin: time.Now(),
		Deleted: false,
	}
}

func (old *Users) ModifyUser(
	username string,
	password string,
	phoneNumber int64,
	) Users {

	//fmt.Println("old user = ", old)

	return Users{
		Id:          old.Id,
		Username:    username,
		Password:    password,
		PhoneNumber: phoneNumber,
		CreatedAt:   old.CreatedAt,
		LastLogin:   time.Now(),
		Deleted:     old.Deleted,
	}
}