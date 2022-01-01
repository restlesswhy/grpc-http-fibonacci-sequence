package httpdel

import (
	"net/http"
	// "net/http"

	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/models"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
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
		var interval models.Interval

		if err := c.Bind(&interval); err != nil {

			return err
		}

		seq, _ := h.fibUC.GetSeq(c.Request().Context(), interval.From, interval.To)
		logger.Info(seq)

		return c.JSON(http.StatusOK, seq)
	}
}