package repository

import (
	"github.com/labstack/gommon/log"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	repoUser "github.com/w33h/Productivity-Tracker-API/repository/users"
	"github.com/w33h/Productivity-Tracker-API/util"
)

type repository struct {
	UserRepository user.RepositoryUser
}

func FactoryRepository(dbCon *util.DatabaseConfig) repository {
	var repo repository

	switch dbCon.Driver {
	case util.Postgres:
		repo.UserRepository = repoUser.NewUserRepository(dbCon.PostgreSQL)
	default:
		log.Info("Unsupported database connection")
	}

	return repo
}