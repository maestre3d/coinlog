package http

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/maestre3d/coinlog/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewCriteria(t *testing.T) {
	e := echo.New()
	baseReq := httptest.NewRequest("GET", "/", nil)
	baseRec := httptest.NewRecorder()
	tests := []struct {
		name     string
		inGenCtx func() echo.Context
		exp      storage.Criteria
	}{
		{
			name: "empty ctx",
			inGenCtx: func() echo.Context {
				return e.NewContext(baseReq, baseRec)
			},
			exp: storage.Criteria{
				Limit:     10,
				PageToken: nil,
			},
		},
		{
			name: "zero limit",
			inGenCtx: func() echo.Context {
				ctx := e.NewContext(baseReq, baseRec)
				ctx.QueryParams().Set("limit", "0")
				return ctx
			},
			exp: storage.Criteria{
				Limit:     10,
				PageToken: nil,
			},
		},
		{
			name: "negative limit",
			inGenCtx: func() echo.Context {
				ctx := e.NewContext(baseReq, baseRec)
				ctx.QueryParams().Set("limit", "-1")
				return ctx
			},
			exp: storage.Criteria{
				Limit:     10,
				PageToken: nil,
			},
		},
		{
			name: "valid limit",
			inGenCtx: func() echo.Context {
				ctx := e.NewContext(baseReq, baseRec)
				ctx.QueryParams().Set("limit", "250")
				return ctx
			},
			exp: storage.Criteria{
				Limit:     250,
				PageToken: nil,
			},
		},
		{
			name: "valid",
			inGenCtx: func() echo.Context {
				ctx := e.NewContext(baseReq, baseRec)
				ctx.QueryParams().Set("limit", "500")
				ctx.QueryParams().Set("page_token", "MQ==")
				return ctx
			},
			exp: storage.Criteria{
				Limit:     500,
				PageToken: storage.NewPageToken("MQ=="),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := newCriteria(tt.inGenCtx())
			assert.EqualValues(t, tt.exp, out)
		})
	}
}
