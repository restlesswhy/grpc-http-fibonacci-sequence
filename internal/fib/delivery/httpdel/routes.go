package httpdel

import (
	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
)

func MapRoutes(fibGroup *echo.Group, h fib.Handler) {
	fibGroup.GET("/seq", h.Get())
}