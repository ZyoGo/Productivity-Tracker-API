package notes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/w33h/Productivity-Tracker-API/api/middleware"
	"github.com/w33h/Productivity-Tracker-API/api/v1/notes/request"
	domain "github.com/w33h/Productivity-Tracker-API/business/notes"
	"github.com/w33h/Productivity-Tracker-API/exception"
	f "github.com/w33h/formatter-response"
)

type Controller struct {
	service domain.ServiceNotes
}

func NewController(service domain.ServiceNotes) *Controller {
	return &Controller{service}
}

func (controller *Controller) CreateNotes(c echo.Context) error {
	userId, _ := middleware.ExtractToken(c)
	createNotesRequest := new(request.CreateRequestNotes)
	if err := c.Bind(createNotesRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *createNotesRequest.ToSpecNotes()

	id, err := controller.service.CreateNote(req, userId)
	if err != nil {
		if err == exception.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		} else if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	message := fmt.Sprintf("Success create Notes with id %v", id)

	return c.JSON(http.StatusCreated, f.CreatedResponse(message))
}

func (controller *Controller) UpdateTodo(c echo.Context) error {
	userId, _ := middleware.ExtractToken(c)
	createNotesRequest := new(request.CreateRequestNotes)
	if err := c.Bind(createNotesRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err))
	}

	id := c.Param("id")
	req := *createNotesRequest.ToSpecNotes()
	req.UserId = userId

	err := controller.service.UpdateNote(req, id)
	if err != nil {
		if err == exception.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		} else if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	message := fmt.Sprintf("Success update Notes with id %v", id)

	return c.JSON(http.StatusOK, f.SuccessResponse(message))
}

func (controller *Controller) DeleteNotes(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("test from notes deleted")

	err := controller.service.DeleteNote(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err))
	}

	message := fmt.Sprintf("Success delete Notes with id %v", id)

	return c.JSON(http.StatusOK, f.SuccessResponse(message))
}

func (controller *Controller) GetNotesByStatus(c echo.Context) error {
	userId, _ := middleware.ExtractToken(c)
	status := c.QueryParam("status")
	fmt.Println("test from notes status")

	notes, err := controller.service.GetNotesByStatus(status, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err))
	}

	return c.JSON(http.StatusOK, f.SuccessResponse(notes))
}

func (controller *Controller) GetNotesById(c echo.Context) error {
	id := c.Param("id")

	notes, err := controller.service.GetNotesById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err))
	}

	return c.JSON(http.StatusOK, f.SuccessResponse(notes))
}

func (controller *Controller) GetNotesByTag(c echo.Context) error {
	userId, _ := middleware.ExtractToken(c)
	tags := c.QueryParam("tags")
	fmt.Println("test from notes tags")

	notes, err := controller.service.GetNotesByTags(tags, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err))
	}

	return c.JSON(http.StatusOK, f.SuccessResponse(notes))
}

func (controller *Controller) GetAllNotes(c echo.Context) error {
	userId, _ := middleware.ExtractToken(c)
	fmt.Println("test from get all notes")
	notes, err := controller.service.GetAllNotes(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err))
	}

	return c.JSON(http.StatusOK, f.SuccessResponse(notes))
}
