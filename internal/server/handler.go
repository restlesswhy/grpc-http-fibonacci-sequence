package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/httpdel"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
)

func (s *Server) MapHandlers(e *echo.Echo, handler fib.Handler) error {

	api := e.Group("/api")

	health := api.Group("/health")

	httpdel.MapRoutes(api, handler)

	health.GET("", func(c echo.Context) error {
		logger.Infof("Health check RequestID: %s", GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}

func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}