package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHealthcheckController(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/health")

	ctrl := NewHealthcheckController()
	require.NoError(t, ctrl.checkLive(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "{\"message\":\"OK\"}\n", rec.Body.String())
}
