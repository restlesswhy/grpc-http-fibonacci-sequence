package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
)

const (
	maxHeaderBytes = 1 << 20
)

type Server struct {
	cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run() error {
	httpserver := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	ctx := context.Background()
	
	router := echo.New()
	router.GET("/fib", getSequence)


	go func() {
		logger.Infof("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := router.StartServer(httpserver); err != nil {
			logger.Fatalf("Error starting Server: ", err)
		}
	}()

	l, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
	select {
	case v := <-quit:
		logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		logger.Errorf("ctx.Done: %v", done)
	}

	if err := router.Shutdown(ctx); err != nil {
		logger.Errorf("Metrics router.Shutdown: %v", err)
	}
	server.GracefulStop()
	logger.Info("Server Exited Properly")

	return nil
}