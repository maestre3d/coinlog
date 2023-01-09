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

type ContactHTTP struct {
	svc appservice.Contact
}

var _ HTTP = ContactHTTP{}

func NewContactHTTP(svc appservice.Contact) ContactHTTP {
	return ContactHTTP{svc: svc}
}

func (u ContactHTTP) MapEndpoints(_ *echo.Echo) {}

func (u ContactHTTP) MapVersionedEndpoints(e *echo.Group) {
	e.POST("/users/:user_id/contacts", u.create)
	e.GET("/users/:user_id/contacts", u.search)
	e.PUT("/users/-/contacts/:contact_id", u.update)
	e.GET("/users/-/contacts/:contact_id", u.getByID)
	e.DELETE("/users/-/contacts/:contact_id", u.deleteByID)
}

func (u ContactHTTP) create(c echo.Context) error {
	type bodyMessage struct {
		DisplayName string `json:"display_name"`
		LinkedTo    string `json:"linked_to"`
		ImageURL    string `json:"image_url"`
	}
	body := bodyMessage{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	id, err := domainservice.KSUIDFactory()
	if err != nil {
		return err
	}
	if err = u.svc.Create(c.Request().Context(), entity.ContactArgs{
		ID:          id,
		DisplayName: body.DisplayName,
		UserID:      c.Param("user_id"),
		LinkedToID:  body.LinkedTo,
		ImageURL:    body.ImageURL,
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, view.BasicResponse{
		Message: id,
	})
}

func (u ContactHTTP) update(c echo.Context) error {
	type bodyMessage struct {
		DisplayName string `json:"display_name"`
		LinkedTo    string `json:"linked_to"`
		ImageURL    string `json:"image_url"`
	}
	body := bodyMessage{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	id := c.Param("contact_id")
	if err := u.svc.Update(c.Request().Context(), entity.ContactArgs{
		ID:          id,
		DisplayName: body.DisplayName,
		UserID:      c.Param("user_id"),
		LinkedToID:  body.LinkedTo,
		ImageURL:    body.ImageURL,
	}); err != nil {
		return err
	}
	return c.JSON(http.StatusAccepted, view.BasicResponse{
		Message: id,
	})
}

func (u ContactHTTP) search(c echo.Context) error {
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

func (u ContactHTTP) getByID(c echo.Context) error {
	usr, err := u.svc.GetByID(c.Request().Context(), c.Param("contact_id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, view.DataResponse{
		Data: usr,
	})
}

func (u ContactHTTP) deleteByID(c echo.Context) error {
	err := u.svc.DeleteByID(c.Request().Context(), c.Param("contact_id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, view.BasicResponse{
		Message: http.StatusText(http.StatusOK),
	})
}
