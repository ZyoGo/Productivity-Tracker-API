package repository

import (
	"github.com/labstack/gommon/log"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	repo "github.com/w33h/Productivity-Tracker-API/repository/users"
	"github.com/w33h/Productivity-Tracker-API/util"
)

type repository struct {
	userRepository user.RepositoryUser
}

func RepositoryFactory(dbCon *util.DatabaseConfig) repository {
	var repository repository

	switch dbCon.Driver {
	case util.Postgres:
		repository.userRepository = repo.NewUserRepository(dbCon.PostgreSQL)
	default:
		log.Info("Unsupported database connection")
	}

	return repository
}