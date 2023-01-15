package http

import (
	"strconv"

	"github.com/maestre3d/coinlog/storage"

	"github.com/labstack/echo/v4"
)

func newCriteria(c echo.Context) storage.Criteria {
	lim, _ := strconv.Atoi(c.QueryParam("limit"))
	return storage.Criteria{
		Limit:     lim,
		PageToken: storage.PageToken(c.QueryParam("page_token")),
	}
}
