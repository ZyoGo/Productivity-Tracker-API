package modules

import (
	"github.com/w33h/Productivity-Tracker-API/api"
	notesV1Controller "github.com/w33h/Productivity-Tracker-API/api/v1/notes"
	todoV1Controller "github.com/w33h/Productivity-Tracker-API/api/v1/todo"
	userV1Controller "github.com/w33h/Productivity-Tracker-API/api/v1/user"
	authController "github.com/w33h/Productivity-Tracker-API/api/v1/auth"
	"github.com/w33h/Productivity-Tracker-API/business/auth"
	"github.com/w33h/Productivity-Tracker-API/business/notes"
	"github.com/w33h/Productivity-Tracker-API/business/todos"
	user "github.com/w33h/Productivity-Tracker-API/business/users"
	dbFactory "github.com/w33h/Productivity-Tracker-API/repository"
	"github.com/w33h/Productivity-Tracker-API/util"
)

func RegisterModules(dbCon *util.DatabaseConfig) api.Controller {
	//Database factory for repository
	dbFactory := dbFactory.FactoryRepository(dbCon)
	// DI for feature User
	userPermitService := user.NewUserService(dbFactory.UserRepository)
	userV1PermitController := userV1Controller.NewController(userPermitService)

	//DI for feature Todos
	todoPermitService := todos.NewTodoService(dbFactory.TodoRepository, dbFactory.UserRepository)
	todoV1PermitController := todoV1Controller.NewController(todoPermitService)

	//DI for feature Notes
	notesPermitService := notes.NewNoteService(dbFactory.NotesRepository, dbFactory.UserRepository)
	notesV1PermitController := notesV1Controller.NewController(notesPermitService)

	//DI for feature Auth
	authPermitService := auth.NewAuthService(dbFactory.AuthRepository)
	authPermitController := authController.NewController(authPermitService)

	controllers := api.Controller{
		UserV1Controller: userV1PermitController,
		TodoV1Controller: todoV1PermitController,
		NotesV1Controller: notesV1PermitController,
		AuthController: authPermitController,
	}

	return controllers
}
