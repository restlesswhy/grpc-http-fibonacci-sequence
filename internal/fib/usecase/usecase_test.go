package usecase

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/mock"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib/models"
	"github.com/stretchr/testify/require"
)

func TestFibUC_GetSeqWithoutCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	cfg := &config.Config{
		Redis: config.RedisConfig{
			Caching: false,
		},
	}

	mockRedis := mock.NewMockRedisRepository(ctrl)
	fibUC := NewFibUC(cfg, mockRedis)

	reqTest := models.Interval{
		From: 2,
		To: 5,
	}

	res, err := fibUC.GetSeq(ctx, reqTest.From, reqTest.To)

	expRes := models.FibSeq{
		Seq: make(map[int32]string),
	}
	expRes.Seq[2] = "1"
	expRes.Seq[3] = "2"
	expRes.Seq[4] = "3"
	expRes.Seq[5] = "5"

	require.NoError(t, err)
	require.Equal(t, expRes, res)
}

func TestFibUC_GetSeqWithCache(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	cfg := &config.Config{
		Redis: config.RedisConfig{
			Caching: true,
		},
	}

	mockRedis := mock.NewMockRedisRepository(ctrl)
	fibUC := NewFibUC(cfg, mockRedis)

	reqTest := models.Interval{
		From: 2,
		To: 5,
	}

	mockRedis.EXPECT().CheckFib(gomock.Any(), gomock.Any()).Return("1", true, nil)
	mockRedis.EXPECT().CheckFib(gomock.Any(), gomock.Any()).Return("2", true, nil)
	mockRedis.EXPECT().CheckFib(gomock.Any(), gomock.Any()).Return("3", true, nil)
	mockRedis.EXPECT().CheckFib(gomock.Any(), gomock.Any()).Return("5", true, nil)
	res, err := fibUC.GetSeq(ctx, reqTest.From, reqTest.To)

	expRes := models.FibSeq{
		Seq: make(map[int32]string),
	}
	expRes.Seq[2] = "1"
	expRes.Seq[3] = "2"
	expRes.Seq[4] = "3"
	expRes.Seq[5] = "5"

	require.NoError(t, err)
	require.Equal(t, expRes, res)
}

func TestFibUC_GetSeqWithCache_2(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	cfg := &config.Config{
		Redis: config.RedisConfig{
			Caching: true,
		},
	}

	mockRedis := mock.NewMockRedisRepository(ctrl)
	fibUC := NewFibUC(cfg, mockRedis)

	reqTest := models.Interval{
		From: 2,
		To: 5,
	}

	mockRedis.EXPECT().CheckFib(gomock.Any(), gomock.Any()).Return("", false, nil)
	mockRedis.EXPECT().Add(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	mockRedis.EXPECT().CheckFib(gomock.Any(), gomock.Any()).Return("2", true, nil)
	mockRedis.EXPECT().CheckFib(gomock.Any(), gomock.Any()).Return("3", true, nil)
	mockRedis.EXPECT().CheckFib(gomock.Any(), gomock.Any()).Return("5", true, nil)
	res, err := fibUC.GetSeq(ctx, reqTest.From, reqTest.To)

	expRes := models.FibSeq{
		Seq: make(map[int32]string),
	}
	expRes.Seq[2] = "1"
	expRes.Seq[3] = "2"
	expRes.Seq[4] = "3"
	expRes.Seq[5] = "5"

	require.NoError(t, err)
	require.Equal(t, expRes, res)
}