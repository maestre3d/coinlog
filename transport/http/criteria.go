package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/domain"
)

func newCriteria(c echo.Context) domain.Criteria {
	lim, _ := strconv.Atoi(c.QueryParam("limit"))
	return domain.Criteria{
		Limit:     lim,
		PageToken: domain.PageToken(c.QueryParam("page_token")),
	}
}
