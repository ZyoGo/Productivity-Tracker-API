package api

import (
	"github.com/labstack/echo/v4"
	"github.com/w33h/Productivity-Tracker-API/api/v1/todo"
	"github.com/w33h/Productivity-Tracker-API/api/v1/user"
)

type Controller struct {
	UserV1Controller *user.Controller
	TodoV1Controller *todo.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	userV1 := e.Group("/v1/user")
	userV1.GET("/:id", controller.UserV1Controller.GetUserById)
	userV1.POST("", controller.UserV1Controller.CreateUser)
	userV1.PUT("/:id", controller.UserV1Controller.UpdateUser)
	userV1.DELETE("/:id", controller.UserV1Controller.DeleteUser)

	todoV1 := e.Group("/v1/todo")
	todoV1.POST("", controller.TodoV1Controller.CreateTodo)
	todoV1.PUT("/:id", controller.TodoV1Controller.UpdateTodo)
	todoV1.DELETE("/:id", controller.TodoV1Controller.DeleteTodo)
	todoV1.GET("", controller.TodoV1Controller.GetByStatus)
	todoV1.GET("/:id", controller.TodoV1Controller.GetById)
	todoV1.GET("", controller.TodoV1Controller.GetAllTodo)
}
