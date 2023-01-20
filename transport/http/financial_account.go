package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/domain/financialaccount"
	"github.com/maestre3d/coinlog/identifier"
	"github.com/maestre3d/coinlog/transport"
)

type FinancialAccountController struct {
	svc financialaccount.Service
}

var _ Controller = FinancialAccountController{}

func NewFinancialController(svc financialaccount.Service) FinancialAccountController {
	return FinancialAccountController{
		svc: svc,
	}
}

func (f FinancialAccountController) MapRoutes(_ *echo.Echo) {
}

func (f FinancialAccountController) MapVersionedRoutes(g *echo.Group) {
	g.POST("/users/:user_id/financial/accounts", f.create)
	g.GET("/users/:user_id/financial/accounts", f.list)
	g.PUT("/users/-/financial/accounts/:account_id", f.update)
	g.GET("/users/-/financial/accounts/:account_id", f.getByID)
	g.DELETE("/users/-/financial/accounts/:account_id", f.delete)
}

func (f FinancialAccountController) create(c echo.Context) error {
	args := financialaccount.CreateCommand{}
	if err := c.Bind(&args); err != nil {
		return err
	}
	id, err := identifier.NewKSUID()
	if err != nil {
		return err
	}
	args.AccountID = id
	args.OwnerID = c.Param("user_id")

	if err = f.svc.Create(c.Request().Context(), args); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, transport.BasicResponse{
		Message: id,
	})
}

func (f FinancialAccountController) update(c echo.Context) error {
	args := financialaccount.UpdateCommand{}
	if err := c.Bind(&args); err != nil {
		return err
	}
	args.AccountID = c.Param("account_id")
	if err := f.svc.Update(c.Request().Context(), args); err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, transport.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (f FinancialAccountController) delete(c echo.Context) error {
	id := c.Param("account_id")
	if err := f.svc.Delete(c.Request().Context(), id); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, transport.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}

func (f FinancialAccountController) getByID(c echo.Context) error {
	acc, err := f.svc.GetByID(c.Request().Context(), c.Param("account_id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, acc)
}

func (f FinancialAccountController) list(c echo.Context) error {
	ls, nextPage, err := f.svc.List(c.Request().Context(), newCriteria(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, transport.ListDataResponse{
		Data:          ls,
		Count:         len(ls),
		NextPageToken: nextPage.String(),
	})
}
