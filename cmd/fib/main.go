package main

import (
	"os"

	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/server"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/logger"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/redis"
)

func main() {
	logger.Info("Starting server...")

	configPath := config.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		logger.Fatalf("cant get config: %v", err)
	}

	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()
	logger.Info("Redis connected")

	s := server.NewServer(cfg, redisClient)
	

	logger.Fatal(s.Run())
}