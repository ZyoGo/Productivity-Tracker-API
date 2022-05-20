package repository

import (
	"github.com/labstack/gommon/log"
	"github.com/w33h/Productivity-Tracker-API/business/auth"
	"github.com/w33h/Productivity-Tracker-API/business/notes"
	"github.com/w33h/Productivity-Tracker-API/business/todos"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	repoNotes "github.com/w33h/Productivity-Tracker-API/repository/notes"
	repoTodo "github.com/w33h/Productivity-Tracker-API/repository/todos"
	repoUser "github.com/w33h/Productivity-Tracker-API/repository/users"
	repoAuth "github.com/w33h/Productivity-Tracker-API/repository/auth"
	"github.com/w33h/Productivity-Tracker-API/util"
)

type repository struct {
	UserRepository user.RepositoryUser
	TodoRepository todos.RepositoryTodos
	NotesRepository notes.RepositoryNotes
	AuthRepository auth.RepositoryAuth
}

func FactoryRepository(dbCon *util.DatabaseConfig) repository {
	var repo repository

	switch dbCon.Driver {
	case util.Postgres:
		repo.UserRepository = repoUser.NewUserRepository(dbCon.PostgreSQL)
		repo.TodoRepository = repoTodo.NewTodoRepository(dbCon.PostgreSQL)
		repo.NotesRepository = repoNotes.NewNotesRepository(dbCon.PostgreSQL)
		repo.AuthRepository = repoAuth.NewAuthRepository(dbCon.PostgreSQL)
	default:
		log.Info("Unsupported database connection")
	}

	return repo
}