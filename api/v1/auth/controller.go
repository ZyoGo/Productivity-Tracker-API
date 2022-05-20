package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/w33h/Productivity-Tracker-API/api/v1/auth/request"
	"github.com/w33h/Productivity-Tracker-API/business/auth"
	"github.com/w33h/Productivity-Tracker-API/exception"
	f "github.com/w33h/formatter-response"
	"net/http"
)

type Controller struct {
	service auth.ServiceAuth
}

func NewController(service auth.ServiceAuth) *Controller {
	return &Controller{service}
}

func (controller *Controller) Auth(c echo.Context) error {
	authRequest := new(request.AuthRequest)
	if err := c.Bind(authRequest); err != nil {
		return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
	}

	req := *authRequest.ToSpec()

	token, err := controller.service.LoginUser(req)
	if err != nil {
		if err == exception.ErrInvalidSpec {
			return c.JSON(http.StatusBadRequest, f.BadRequestResponse(err.Error()))
		}else if err == exception.ErrNotFound {
			return c.JSON(http.StatusNotFound, f.NotFoundResponse(err.Error()))
		}
		return c.JSON(http.StatusInternalServerError, f.InternalServerErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusCreated, f.CreatedResponse(token))
}