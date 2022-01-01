package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/httpdel"
)

func (s *Server) MapHandlers(e *echo.Echo, handler fib.Handler) error {

	api := e.Group("/api")

	health := api.Group("/health")

	httpdel.MapRoutes(api, handler)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
