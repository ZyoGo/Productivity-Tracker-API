package modules

import (
	"github.com/w33h/Productivity-Tracker-API/api"
	userV1Controller "github.com/w33h/Productivity-Tracker-API/api/v1/user"
	domain "github.com/w33h/Productivity-Tracker-API/business/users"
	"github.com/w33h/Productivity-Tracker-API/config"
	userRepo "github.com/w33h/Productivity-Tracker-API/repository"
	"github.com/w33h/Productivity-Tracker-API/util"
)

func RegisterModules(dbCon *util.DatabaseConfig, config *config.AppConfig) api.Controller {
	userPermitRepository := userRepo.FactoryRepository(dbCon)
	userPermitService := domain.NewUserService(userPermitRepository.UserRepository)

	userV1PermitController := userV1Controller.NewController(userPermitService)

	controllers := api.Controller{
		UserV1Controller: userV1PermitController,
	}

	return controllers
}