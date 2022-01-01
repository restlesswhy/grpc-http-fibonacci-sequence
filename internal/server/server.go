package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/grpcdel"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/httpdel"
	pb "github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/proto"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/repository"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/usecase"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	maxHeaderBytes = 1 << 20
)

type Server struct {
	echo        *echo.Echo
	cfg *config.Config
	redisClient *redis.Client
}

func NewServer(cfg *config.Config, redisClient *redis.Client) *Server {
	return &Server{
		echo: echo.New(),
		cfg: cfg,
		redisClient: redisClient,
	}
}

func (s *Server) Run() error {
	httpserver := &http.Server{
		Addr:           s.cfg.ServerHttp.Port,
		ReadTimeout:    time.Second * s.cfg.ServerHttp.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.ServerHttp.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	ctx := context.Background()
	redisRepo := repository.NewRedisRepo(s.redisClient, s.cfg)
	fibUC := usecase.NewFibUC(s.cfg, redisRepo)
	fibHandler := httpdel.NewFibHandler(fibUC)

	if err := s.MapHandlers(s.echo, fibHandler); err != nil {
		return err
	}

	go func() {
		logger.Infof("HTTP server is listening on PORT: %s", s.cfg.ServerHttp.Port)
		if err := s.echo.StartServer(httpserver); err != nil {
			logger.Fatalf("Error starting Server: ", err)
		}
	}()

	l, err := net.Listen("tcp", s.cfg.ServerGrpc.Port)
	if err != nil {
		return err
	}
	defer l.Close()

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.ServerGrpc.MaxConnectionIdle * time.Minute,
		Timeout:           s.cfg.ServerGrpc.Timeout * time.Second,
		MaxConnectionAge:  s.cfg.ServerGrpc.MaxConnectionAge * time.Minute,
		Time:              s.cfg.ServerGrpc.Timeout * time.Minute,
	}))

	fiboGrpcMicroservice := grpcdel.NewFibMicroservice(fibUC)
	pb.RegisterFiboSequenceServiceServer(server, fiboGrpcMicroservice)

	go func() {
		logger.Infof("gRPC server is listening on port: %v", s.cfg.ServerGrpc.Port)
		logger.Fatal(server.Serve(l))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	
	select {
	case v := <-quit:
		logger.Errorf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		logger.Errorf("ctx.Done: %v", done)
	}

	if err := s.echo.Shutdown(ctx); err != nil {
		logger.Errorf("router.Shutdown: %v", err)
	}
	server.GracefulStop()
	logger.Info("Servers Exited Properly")

	return nil
}

