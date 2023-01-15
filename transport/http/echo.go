package http

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-multierror"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maestre3d/coinlog/exception"
	"github.com/maestre3d/coinlog/transport"
	"github.com/rs/zerolog/log"
	"golang.org/x/time/rate"
)

// NewEcho builds a pre-configured and ready-to-use echo.Echo instance.
func NewEcho(cfg Config, mapper *ControllerMapper) *echo.Echo {
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
	mapper.RegisterRoutes(e)
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
	var res transport.ErrorsResponse
	if echoErr, ok := err.(*echo.HTTPError); ok {
		var msg string
		switch v := echoErr.Message.(type) {
		case string:
			msg = v
		}
		res = transport.ErrorsResponse{
			Errors: []transport.ErrorResponse{
				{
					Code:        echoErr.Code,
					ErrorStatus: msg,
					Message:     msg,
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
func newErrorsMessage(err error) (res transport.ErrorsResponse) {
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
		res.Errors = make([]transport.ErrorResponse, 0, len(errs))
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
	res.Errors = []transport.ErrorResponse{errMsg}
	return
}

// newErrorMessage builds a ErrorMessage from a single error (no multierror.Error).
func newErrorMessage(err error) transport.ErrorResponse {
	switch v := err.(type) {
	case validator.FieldError:
		return newErrorMessage(exception.NewFromValidator(v))
	case exception.ResourceNotFound:
		return transport.ErrorResponse{
			Message:     v.Error(),
			Code:        http.StatusNotFound,
			ErrorStatus: v.TypeName(),
		}
	case exception.DomainGeneric:
		return transport.ErrorResponse{
			Message:     v.Error(),
			Code:        http.StatusBadRequest,
			ErrorStatus: v.TypeName(),
		}
	case exception.MissingParameter:
		return transport.ErrorResponse{
			Message:     v.Error(),
			Code:        http.StatusBadRequest,
			ErrorStatus: v.TypeName(),
		}
	case exception.ParameterOutOfRange:
		return transport.ErrorResponse{
			Message:     v.Error(),
			Code:        http.StatusBadRequest,
			ErrorStatus: v.TypeName(),
		}
	case exception.InvalidParameter:
		return transport.ErrorResponse{
			Message:     v.Error(),
			Code:        http.StatusBadRequest,
			ErrorStatus: v.TypeName(),
		}
	case exception.UnhealthyNode:
		return transport.ErrorResponse{
			Code:        http.StatusServiceUnavailable,
			ErrorStatus: v.TypeName(),
			Message:     v.Error(),
		}
	default:
		return transport.ErrorResponse{
			Message:     "error",
			Code:        http.StatusInternalServerError,
			ErrorStatus: http.StatusText(http.StatusInternalServerError),
		}
	}
}
