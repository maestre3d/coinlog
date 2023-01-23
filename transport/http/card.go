package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/domain/card"
	"github.com/maestre3d/coinlog/identifier"
	"github.com/maestre3d/coinlog/transport"
)

type CardController struct {
	svc card.Service
}

var _ Controller = CardController{}

func NewCardController(s card.Service) CardController {
	return CardController{
		svc: s,
	}
}

func (u CardController) MapRoutes(_ *echo.Echo) {}

func (u CardController) MapVersionedRoutes(g *echo.Group) {
	g.POST("/users/:user_id/cards", u.create)
	g.GET("/users/:user_id/cards", u.list)
	g.PUT("/users/-/cards/:card_id", u.update)
	g.GET("/users/-/cards/:card_id", u.getByID)
	g.DELETE("/users/-/cards/:card_id", u.delete)
}

func (u CardController) create(c echo.Context) error {
	args := card.CreateCommand{}
	if err := c.Bind(&args); err != nil {
		return err
	}
	args.OwnerID = c.Param("user_id")

	id, err := identifier.NewKSUID()
	if err != nil {
		return err
	}

	args.CardID = id
	if err = u.svc.Create(c.Request().Context(), args); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, transport.BasicResponse{
		Message: id,
	})
}

func (u CardController) update(c echo.Context) error {
	args := card.UpdateCommand{}
	if err := c.Bind(&args); err != nil {
		return err
	}
	args.CardID = c.Param("card_id")
	if err := u.svc.Update(c.Request().Context(), args); err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, transport.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (u CardController) delete(c echo.Context) error {
	id := c.Param("card_id")
	if err := u.svc.Delete(c.Request().Context(), id); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transport.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (u CardController) getByID(c echo.Context) error {
	id := c.Param("card_id")
	usr, err := u.svc.GetByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transport.DataResponse{
		Data: usr,
	})
}

func (u CardController) list(c echo.Context) error {
	ls, nextPage, err := u.svc.ListUserCards(c.Request().Context(), newCriteria(c), c.Param("user_id"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transport.ListDataResponse{
		Data:          ls,
		Count:         len(ls),
		NextPageToken: nextPage.String(),
	})
}
