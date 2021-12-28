package main

import (
	"os"

	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/server"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
)

func main() {
	logger.Info("Starting server...")

	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		logger.Fatalf("cant get config: %v", err)
	}

	s := server.NewServer(cfg)
	
	logger.Fatal(s.Run())
}