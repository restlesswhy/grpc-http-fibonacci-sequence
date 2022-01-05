package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/pkg/errors"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/grpcdel"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/httpdel"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/repository"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/usecase"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/pkg/redis"
	"github.com/stretchr/testify/suite"
)

type FibTestSuite struct {
	suite.Suite

	cfg *config.Config
	cache *cache.Cache
	httpHandler *httpdel.FibHandler
	grpcHandler *grpcdel.FibMicroservice
	usecase *usecase.FibUC
	repo *repository.RedisRepo
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(FibTestSuite))
}

func (s *FibTestSuite) SetupSuite() {
	cfg := &config.Config{
		Redis: config.RedisConfig{
			RedisAddr: "127.0.0.1:6379",
			FibTTL: time.Second * 5,
		},
	}
	s.cfg = cfg

	redisClient := redis.NewRedisClient(cfg)
	
	cache := cache.New(&cache.Options{
        Redis:      redisClient,
        LocalCache: cache.NewTinyLFU(1000, s.cfg.Redis.FibTTL),
    })

	s.cache = cache

	s.initDeps()

	if err := s.populateDB(); err != nil {
		s.FailNow("Failed to populate DB", err)
	}
}

func (s *FibTestSuite) TearDownSuite() {
	
}

func (s *FibTestSuite) initDeps() {
	// Init domain deps
	repos := repository.NewRedisRepo(s.cache, s.cfg)
	usecase := usecase.NewFibUC(s.cfg, repos)
	
	s.repo = repos
	s.usecase = usecase
	s.httpHandler = httpdel.NewFibHandler(usecase)
	s.grpcHandler = grpcdel.NewFibMicroservice(usecase)
}

func TestMain(m *testing.M) {
	rc := m.Run()
	os.Exit(rc)
}

func (s *FibTestSuite) populateDB() error {
	if err := s.cache.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   "5",
		Value: "5",
		TTL: s.cfg.Redis.FibTTL,
	}); err != nil {
		return errors.Wrap(err, "tests populate")
	}

	if err := s.cache.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   "6",
		Value: "8",
		TTL: s.cfg.Redis.FibTTL,
	}); err != nil {
		return errors.Wrap(err, "tests populate")
	}

	if err := s.cache.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   "7",
		Value: "13",
		TTL: s.cfg.Redis.FibTTL,
	}); err != nil {
		return errors.Wrap(err, "tests populate")
	}

	return nil
}