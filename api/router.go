package api

import (
	"github.com/labstack/echo/v4"
	"github.com/w33h/Productivity-Tracker-API/api/middleware"
	"github.com/w33h/Productivity-Tracker-API/api/v1/auth"
	"github.com/w33h/Productivity-Tracker-API/api/v1/notes"
	"github.com/w33h/Productivity-Tracker-API/api/v1/todo"
	"github.com/w33h/Productivity-Tracker-API/api/v1/user"
)

type Controller struct {
	UserV1Controller  *user.Controller
	TodoV1Controller  *todo.Controller
	NotesV1Controller *notes.Controller

	AuthController *auth.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.POST("/user", controller.UserV1Controller.CreateUser)

	userV1 := e.Group("/v1/user")
	userV1.Use(middleware.JWTMiddleware())
	userV1.GET("/:id", controller.UserV1Controller.GetUserById)
	userV1.PUT("/:id", controller.UserV1Controller.UpdateUser)
	userV1.DELETE("/:id", controller.UserV1Controller.DeleteUser)

	notesV1 := e.Group("/v1")
	notesV1.Use(middleware.JWTMiddleware())
	notesV1.GET("/notes", controller.NotesV1Controller.GetAllNotes)
	notesV1.GET("/note/:id", controller.NotesV1Controller.GetNotesById)
	notesV1.POST("/note", controller.NotesV1Controller.CreateNotes)
	notesV1.PUT("/note/:id", controller.NotesV1Controller.UpdateTodo)
	notesV1.DELETE("/note/:id", controller.NotesV1Controller.DeleteNotes)
	notesV1.GET("/notes/tag", controller.NotesV1Controller.GetNotesByTag)
	notesV1.GET("/notes/status", controller.NotesV1Controller.GetNotesByStatus)

	todoV1 := e.Group("/v1/todo")
	todoV1.Use(middleware.JWTMiddleware())
	todoV1.POST("", controller.TodoV1Controller.CreateTodo)
	todoV1.PUT("/:id", controller.TodoV1Controller.UpdateTodo)
	todoV1.DELETE("/:id", controller.TodoV1Controller.DeleteTodo)
	todoV1.GET("/status", controller.TodoV1Controller.GetByStatus)
	todoV1.GET("/:id", controller.TodoV1Controller.GetById)
	todoV1.GET("", controller.TodoV1Controller.GetAllTodo)

	auth := e.Group("/v1/auth")
	auth.POST("", controller.AuthController.Auth)
}
