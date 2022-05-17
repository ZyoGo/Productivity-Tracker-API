package user

import (
	"fmt"
	f "github.com/w33h/formatter-response"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/w33h/Productivity-Tracker-API/api/v1/user/request"
	domain "github.com/w33h/Productivity-Tracker-API/business/users"
)

type Controller struct {
	service domain.ServiceUser
}

func NewController(service domain.ServiceUser) *Controller {
	return &Controller{
		service: service,
	}
}

func (controller *Controller) GetUserById(c echo.Context) error {
	user, err := controller.service.GetById(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *Controller) CreateUser(c echo.Context) error {
	createUserRequest := new(request.CreateRequestUser)

	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err))
	}

	req := *createUserRequest.ToSpecUser()

	result, err := controller.service.CreateUser(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	fmt.Println("create user request = ", result)

	return c.JSON(http.StatusCreated, f.CreatedResponse(result))
}

func (controller *Controller) UpdateUser(c echo.Context) error {
	createUserRequest := new(request.CreateRequestUser)
	id := c.Param("id")

	if err := c.Bind(createUserRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err))
	}

	req := *createUserRequest.ToSpecUser()

	err := controller.service.UpdateUser(req, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	message := fmt.Sprintf("Success update user with id %v", id)

	return c.JSON(http.StatusOK, f.SuccessResponse(message))
}

func (controller *Controller) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := controller.service.DeleteUser(id)
	fmt.Println("error message = ", err)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	message := fmt.Sprintf("Success delete user with id %v", id)

	return c.JSON(http.StatusOK, f.SuccessResponse(message))
}