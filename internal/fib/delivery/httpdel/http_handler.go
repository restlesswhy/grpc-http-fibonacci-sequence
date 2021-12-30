package httpdel

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
)

type fibHandler struct {
	fibUC fib.UseCase
}

func NewFibHandler(fibUC fib.UseCase) fib.Handler {
	return &fibHandler{
		fibUC: fibUC,
	}
}

func (h *fibHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})

	}
}