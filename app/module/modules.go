package modules

import (
	"github.com/w33h/Productivity-Tracker-API/api"
	userV1Controller "github.com/w33h/Productivity-Tracker-API/api/v1/user"
	todoV1Controller "github.com/w33h/Productivity-Tracker-API/api/v1/todo"
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

	//DI for feature Todo
	todoPermitService := todos.NewTodoService(dbFactory.TodoRepository, dbFactory.UserRepository)
	todoV1PermitController := todoV1Controller.NewController(todoPermitService)

	controllers := api.Controller{
		UserV1Controller: userV1PermitController,
		TodoV1Controller: todoV1PermitController,
	}

	return controllers
}
