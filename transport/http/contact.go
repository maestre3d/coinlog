package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/domain/contact"
	"github.com/maestre3d/coinlog/identifier"
	"github.com/maestre3d/coinlog/transport"
)

type ContactController struct {
	svc contact.Service
}

var _ Controller = ContactController{}

func NewContactController(s contact.Service) ContactController {
	return ContactController{
		svc: s,
	}
}

func (u ContactController) MapRoutes(_ *echo.Echo) {}

func (u ContactController) MapVersionedRoutes(g *echo.Group) {
	g.POST("/users/:user_id/contacts", u.create)
	g.GET("/users/:user_id/contacts", u.list)
	g.PUT("/users/-/contacts/:contact_id", u.update)
	g.GET("/users/-/contacts/:contact_id", u.getByID)
	g.DELETE("/users/-/contacts/:contact_id", u.delete)
}

func (u ContactController) create(c echo.Context) error {
	args := contact.CreateCommand{}
	if err := c.Bind(&args); err != nil {
		return err
	}
	args.UserID = c.Param("user_id")

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

func (u ContactController) update(c echo.Context) error {
	args := contact.UpdateCommand{}
	if err := c.Bind(&args); err != nil {
		return err
	}
	args.ID = c.Param("contact_id")
	if err := u.svc.Update(c.Request().Context(), args); err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, transport.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (u ContactController) delete(c echo.Context) error {
	id := c.Param("contact_id")
	if err := u.svc.Delete(c.Request().Context(), id); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transport.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (u ContactController) getByID(c echo.Context) error {
	id := c.Param("contact_id")
	usr, err := u.svc.GetByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transport.DataResponse{
		Data: usr,
	})
}

func (u ContactController) list(c echo.Context) error {
	ls, nextPage, err := u.svc.ListUserContacts(c.Request().Context(), newCriteria(c), c.Param("user_id"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transport.ListDataResponse{
		Data:          ls,
		Count:         len(ls),
		NextPageToken: nextPage.String(),
	})
}
