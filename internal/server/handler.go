package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/httpdel"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/repository"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/usecase"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
)

func (s *Server) MapHandlers(e *echo.Echo) error {

	redisRepo := repository.NewRedisRepo(s.redisClient)
	fibUC := usecase.NewFibUC(s.cfg, redisRepo)
	httpHandler := httpdel.NewFibHandler(fibUC)

	api := e.Group("/api")

	health := api.Group("/health")

	httpdel.MapRoutes(api, httpHandler)

	health.GET("", func(c echo.Context) error {
		logger.Infof("Health check RequestID: %s", GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}

func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}