package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/domain/user"
	"github.com/maestre3d/coinlog/identifier"
	"github.com/maestre3d/coinlog/transport"
)

type UserController struct {
	svc user.Service
}

var _ Controller = UserController{}

func NewUserController(s user.Service) UserController {
	return UserController{
		svc: s,
	}
}

func (u UserController) MapRoutes(e *echo.Echo) {}

func (u UserController) MapVersionedRoutes(g *echo.Group) {
	g.POST("/users", u.create)
	g.GET("/users", u.list)
	g.PUT("/users/:user_id", u.update)
	g.GET("/users/:user_id", u.getByID)
	g.DELETE("/users/:user_id", u.delete)
}

func (u UserController) create(c echo.Context) error {
	args := user.CreateArgs{}
	if err := c.Bind(&args); err != nil {
		return err
	}

	id, err := identifier.NewKSUID()
	if err != nil {
		return err
	}

	args.ID = id
	if err = u.svc.Create(c.Request().Context(), args); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, transport.BasicResponse{
		Message: id,
	})
}

func (u UserController) update(c echo.Context) error {
	args := user.UpdateArgs{}
	if err := c.Bind(&args); err != nil {
		return err
	}
	args.ID = c.Param("user_id")
	if err := u.svc.Update(c.Request().Context(), args); err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, transport.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (u UserController) delete(c echo.Context) error {
	id := c.Param("user_id")
	if err := u.svc.Delete(c.Request().Context(), id); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transport.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (u UserController) getByID(c echo.Context) error {
	id := c.Param("user_id")
	usr, err := u.svc.GetByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transport.DataResponse{
		Data: usr,
	})
}

func (u UserController) list(c echo.Context) error {
	ls, nextPage, err := u.svc.List(c.Request().Context(), newCriteria(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transport.ListDataResponse{
		Data:          ls,
		Count:         len(ls),
		NextPageToken: nextPage.String(),
	})
}
