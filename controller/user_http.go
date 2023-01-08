package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/appservice"
	"github.com/maestre3d/coinlog/domainservice"
	"github.com/maestre3d/coinlog/view"
)

type UserHTTP struct {
	svc appservice.User
}

var _ HTTP = UserHTTP{}

func NewUserHTTP(svc appservice.User) UserHTTP {
	return UserHTTP{svc: svc}
}

func (u UserHTTP) MapEndpoints(_ *echo.Echo) {}

func (u UserHTTP) MapVersionedEndpoints(e *echo.Group) {
	e.POST("/users", u.createUser)
}

func (u UserHTTP) createUser(c echo.Context) error {
	body := view.User{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	id, err := domainservice.KSUIDFactory()
	if err != nil {
		return err
	}
	if err = u.svc.CreateUser(c.Request().Context(), id, body.DisplayName); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, view.BasicResponse{
		Message: http.StatusText(http.StatusCreated),
	})
}
