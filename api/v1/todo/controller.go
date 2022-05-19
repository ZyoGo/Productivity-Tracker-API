package todo

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/w33h/Productivity-Tracker-API/api/v1/todo/request"
	domain "github.com/w33h/Productivity-Tracker-API/business/todos"
	"github.com/w33h/Productivity-Tracker-API/exception"
	f "github.com/w33h/formatter-response"
	"net/http"
)

type Controller struct {
	service domain.ServiceTodos
}

func NewController(service domain.ServiceTodos) *Controller {
	return &Controller{service}
}

func (controller *Controller) CreateTodo(c echo.Context) error {
	createTodoRequest := new(request.CreateRequestTodo)
	if err := c.Bind(createTodoRequest); err != nil {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	fmt.Println("createtodorequest = ", createTodoRequest)

	req := *createTodoRequest.ToSpecTodo()
	fmt.Println("req = ", req)

	id, err := controller.service.CreateTodo(req)
	if err != nil {
		if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		} else if err == exception.ErrInvalidSpec{
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	message := fmt.Sprintf("Success create Todo with id %v", id)

	return c.JSON(http.StatusCreated, f.CreatedResponse(message))
}

func (controller *Controller) UpdateTodo(c echo.Context) error {
	createRequestTodo := new(request.CreateRequestTodo)
	if err := c.Bind(createRequestTodo); err != nil {
		return c.JSON(http.StatusBadGateway, f.BadGatewayResponse(err))
	}

	id := c.Param("id")

	err := controller.service.UpdateTodo(*createRequestTodo.ToSpecTodo(),id)
	if err != nil {
		if err == exception.ErrInvalidSpec {
			return c.JSON(http.StatusBadGateway, f.BadGatewayResponse(err))
		}else if err == exception.ErrNotFound{
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err))
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	message := fmt.Sprintf("Success update Todo with id %v", id)

	return c.JSON(http.StatusOK, f.SuccessResponse(message))
}

func (controller *Controller) DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	err := controller.service.DeleteTodo(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, f.NotFoundResponse(err))
	}

	message := fmt.Sprintf("Success delete Todo with id %v", id)

	return c.JSON(http.StatusOK, f.SuccessResponse(message))
}

func (controller *Controller) GetByStatus(c echo.Context) error {
	status := c.QueryParam("status")

	result, err := controller.service.GetByStatus(status)
	if err != nil {
		return c.JSON(http.StatusNotFound, f.NotFoundResponse(err))
	}

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}

func (controller *Controller) GetById(c echo.Context) error {
	id := c.Param("id")

	result, err := controller.service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, f.NotFoundResponse(err))
	}

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}

func (controller *Controller) GetAllTodo(c echo.Context) error {
	result, err := controller.service.GetAllTodo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err))
	}

	return c.JSON(http.StatusOK, f.SuccessResponse(result))
}