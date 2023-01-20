package http

import (
	"strconv"

	"github.com/maestre3d/coinlog/storage"

	"github.com/labstack/echo/v4"
)

func newCriteria(c echo.Context) (out storage.Criteria) {
	out.Limit, _ = strconv.Atoi(c.QueryParam("limit"))
	if out.Limit <= 0 {
		out.Limit = 10
	}

	if token := c.QueryParam("page_token"); token != "" {
		out.PageToken = storage.NewPageToken(token)
	}
	return
}
