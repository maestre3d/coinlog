package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maestre3d/coinlog/configuration"
	"github.com/maestre3d/coinlog/exception"
	"github.com/maestre3d/coinlog/view"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
)

// HTTP contains one or many endpoints to be exposed using the HTTP protocol.
type HTTP interface {
	// MapEndpoints associate internal functions to HTTP actions and routes (e.g. GET /foo).
	MapEndpoints(e *echo.Echo)
	// MapVersionedEndpoints associate internal functions to HTTP actions and versioned routes (e.g. POST /v1/foo).
	MapVersionedEndpoints(e *echo.Group)
}

// NewEcho builds a pre-configured and ready-to-use echo.Echo instance.
func NewEcho(cfg configuration.ServerHTTP, m *MuxHTTP) *echo.Echo {
	e := echo.New()

	// Setup middlewares
	e.HTTPErrorHandler = echoErrHandler
	e.Use(middleware.Recover())
	e.Use(middleware.Timeout())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(cfg.MaxRate)))) // default 100/sec
	e.Use(middleware.BodyLimit(cfg.MaxRequestBodySize))                                          // 20 MB by default
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	//e.Use(middleware.CSRF())
	e.Use(echoLoggerFunc())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	// Setup routers
	m.RegisterRoutes(e)
	return e
}

// echoLoggerFunc builds an echo.MiddlewareFunc which logs every request using zerolog's log.Log.
func echoLoggerFunc() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		Skipper:        middleware.DefaultSkipper,
		BeforeNextFunc: nil,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			contentLength, _ := strconv.ParseInt(v.ContentLength, 10, 64)
			log.Info().
				Str("id", v.RequestID).
				Str("source_ip", v.RemoteIP).
				Str("host", v.Host).
				Str("protocol", v.Protocol).
				Str("method", v.Method).
				Str("uri", v.URI).
				Str("user_agent", v.UserAgent).
				Int("status", v.Status).
				Str("latency_human", v.Latency.String()).
				Int64("content_length", contentLength).
				Int64("response_size", v.ResponseSize).
				Msg("http request received")
			if v.Error == nil {
				return nil
			}
			errs, ok := v.Error.(*multierror.Error)
			if !ok {
				log.Err(v.Error).
					Str("id", v.RequestID).
					Str("source_ip", v.RemoteIP).
					Str("host", v.Host).
					Str("protocol", v.Protocol).
					Str("method", v.Method).
					Str("uri", v.URI).
					Str("user_agent", v.UserAgent).
					Int("status", v.Status).
					Int64("content_length", contentLength).
					Int64("response_size", v.ResponseSize).
					Msg("http request failed")
				return v.Error
			}

			for _, err := range errs.Errors {
				log.Err(err).
					Str("id", v.RequestID).
					Str("source_ip", v.RemoteIP).
					Str("host", v.Host).
					Str("protocol", v.Protocol).
					Str("method", v.Method).
					Str("uri", v.URI).
					Str("user_agent", v.UserAgent).
					Int("status", v.Status).
					Int64("content_length", contentLength).
					Int64("response_size", v.ResponseSize).
					Msg("http request failed")
			}
			return v.Error
		},
		LogProtocol:      true,
		LogLatency:       true,
		LogRemoteIP:      true,
		LogHost:          true,
		LogMethod:        true,
		LogURI:           true,
		LogRequestID:     true,
		LogUserAgent:     true,
		LogStatus:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogError:         true,
	})
}

// echoErrHandler handles errors from an echo.Echo HTTP request.
func echoErrHandler(err error, c echo.Context) {
	var res view.ErrorsMessage
	if echoErr, ok := err.(*echo.HTTPError); ok {
		var msg string
		switch v := echoErr.Message.(type) {
		case string:
			msg = v
		}
		res = view.ErrorsMessage{
			Errors: []view.ErrorMessage{
				{
					Code:      echoErr.Code,
					ErrorType: msg,
					Message:   msg,
				},
			},
			Code: echoErr.Code,
		}
	} else {
		res = newErrorsMessage(err)
	}
	_ = c.JSON(res.Code, res)
}

// newErrorsMessage builds a ErrorsMessage from a given error.
// Might traverse errors slice if multierror.Error is used.
func newErrorsMessage(err error) (res view.ErrorsMessage) {
	var errs []error
	if errMultierr, ok := err.(*multierror.Error); ok {
		errs = errMultierr.Errors
	} else if errValidator, okValidator := err.(validator.ValidationErrors); okValidator {
		errs = make([]error, 0, len(errValidator))
		for _, e := range errValidator {
			errs = append(errs, e)
		}
	}

	if len(errs) > 0 {
		res.Errors = make([]view.ErrorMessage, 0, len(errs))
		for _, errWrap := range errs {
			errMsg := newErrorMessage(errWrap)
			if res.Code < errMsg.Code {
				res.Code = errMsg.Code
			}
			res.Errors = append(res.Errors, errMsg)
		}
		return
	}

	errMsg := newErrorMessage(err)
	res.Code = errMsg.Code
	res.Errors = []view.ErrorMessage{errMsg}
	return
}

// newErrorMessage builds a ErrorMessage from a single error (no multierror.Error).
func newErrorMessage(err error) view.ErrorMessage {
	switch v := err.(type) {
	case validator.FieldError:
		return newErrorMessage(exception.NewFromValidator(v))
	case exception.ResourceNotFound:
		return view.ErrorMessage{
			Message:   v.Error(),
			Code:      http.StatusNotFound,
			ErrorType: v.TypeName(),
		}
	case exception.DomainUnknown:
		return view.ErrorMessage{
			Message:   v.Error(),
			Code:      http.StatusBadRequest,
			ErrorType: v.TypeName(),
		}
	case exception.MissingParameter:
		return view.ErrorMessage{
			Message:   v.Error(),
			Code:      http.StatusBadRequest,
			ErrorType: v.TypeName(),
		}
	case exception.ParameterOutOfRange:
		return view.ErrorMessage{
			Message:   v.Error(),
			Code:      http.StatusBadRequest,
			ErrorType: v.TypeName(),
		}
	case exception.InvalidParameter:
		return view.ErrorMessage{
			Message:   v.Error(),
			Code:      http.StatusBadRequest,
			ErrorType: v.TypeName(),
		}
	case exception.UnhealthyNode:
		return view.ErrorMessage{
			Code:      http.StatusServiceUnavailable,
			ErrorType: v.TypeName(),
			Message:   v.Error(),
		}
	default:
		return view.ErrorMessage{
			Message:   "error",
			Code:      http.StatusInternalServerError,
			ErrorType: http.StatusText(http.StatusInternalServerError),
		}
	}
}
