package repository

import (
	"github.com/labstack/gommon/log"
	"github.com/w33h/Productivity-Tracker-API/business/todos"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	repoTodo "github.com/w33h/Productivity-Tracker-API/repository/todos"
	repoUser "github.com/w33h/Productivity-Tracker-API/repository/users"
	"github.com/w33h/Productivity-Tracker-API/util"
)

type repository struct {
	UserRepository user.RepositoryUser
	TodoRepository todos.RepositoryTodos
}

func FactoryRepository(dbCon *util.DatabaseConfig) repository {
	var repo repository

	switch dbCon.Driver {
	case util.Postgres:
		repo.UserRepository = repoUser.NewUserRepository(dbCon.PostgreSQL)
		repo.TodoRepository = repoTodo.NewTodoRepository(dbCon.PostgreSQL)
	default:
		log.Info("Unsupported database connection")
	}

	return repo
}