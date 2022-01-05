package tests

import (
	"testing"

	"github.com/go-redis/cache/v8"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/grpcdel"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/delivery/httpdel"
	"github.com/stretchr/testify/suite"
)

type FibTestSuite struct {
	suite.Suite

	cache *cache.Cache
	httpHandler *httpdel.FibHandler
	grpcHandler *grpcdel.FibMicroservice
}

func TestAPISuite(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	suite.Run(t, new(FibTestSuite))
}

func (s *APITestSuite) SetupSuite() {
	
}
