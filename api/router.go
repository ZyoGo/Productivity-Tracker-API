package api

import (
	"github.com/labstack/echo/v4"
	"github.com/w33h/Productivity-Tracker-API/api/v1/user"
)

type Controller struct {
	UserV1Controller *user.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	userV1 := e.Group("/v1/user")
	userV1.GET("/:id", controller.UserV1Controller.GetUserById)
	userV1.POST("", controller.UserV1Controller.CreateUser)
	userV1.PUT("/:id", controller.UserV1Controller.UpdateUser)
	userV1.DELETE("/:id", controller.UserV1Controller.DeleteUser)
}
