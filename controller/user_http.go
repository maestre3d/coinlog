package controller

import (
	"github.com/maestre3d/coinlog/entity"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/appservice"
	"github.com/maestre3d/coinlog/domainservice"
	"github.com/maestre3d/coinlog/valueobject"
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
	e.POST("/users", u.create)
	e.GET("/users", u.search)
	e.PUT("/users/:user_id", u.update)
	e.GET("/users/:user_id", u.getByID)
	e.DELETE("/users/:user_id", u.deleteByID)
}

func (u UserHTTP) create(c echo.Context) error {
	body := view.User{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	id, err := domainservice.KSUIDFactory()
	if err != nil {
		return err
	}
	if err = u.svc.Create(c.Request().Context(), entity.UserArgs{
		ID:          id,
		DisplayName: body.DisplayName,
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, view.BasicResponse{
		Message: id,
	})
}

func (u UserHTTP) update(c echo.Context) error {
	body := view.User{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	id := c.Param("user_id")
	if err := u.svc.Update(c.Request().Context(), entity.UserArgs{
		ID:          id,
		DisplayName: body.DisplayName,
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, view.BasicResponse{
		Message: id,
	})
}

func (u UserHTTP) search(c echo.Context) error {
	lim, _ := strconv.Atoi(c.QueryParam("limit"))
	usr, token, err := u.svc.Search(c.Request().Context(), valueobject.Criteria{
		Limit:     lim,
		PageToken: valueobject.PageToken(c.QueryParam("page_token")),
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, view.ListDataResponse{
		Items:         usr,
		Count:         len(usr),
		NextPageToken: token.String(),
	})
}

func (u UserHTTP) getByID(c echo.Context) error {
	usr, err := u.svc.GetByID(c.Request().Context(), c.Param("user_id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, view.DataResponse{
		Data: usr,
	})
}

func (u UserHTTP) deleteByID(c echo.Context) error {
	err := u.svc.DeleteByID(c.Request().Context(), c.Param("user_id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, view.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}
