package repository

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/internal/fib"
	"github.com/stretchr/testify/require"
)

func SetupRedis() fib.RedisRepository {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatal(err)
	}
	client := redis.NewRing(&redis.RingOptions{
        Addrs: map[string]string{
            "server1": mr.Addr(),
        },
    })

	cache := cache.New(&cache.Options{
        Redis:      client,
        LocalCache: cache.NewTinyLFU(1000, time.Second * 2),
    })

	cfg := config.Config{
		Redis: config.RedisConfig{
			FibTTL: time.Second * 2,
		},
	}

	fibRedisRepo := NewRedisRepo(cache, &cfg)
	return fibRedisRepo
}
func TestRedisRepo_Add(t *testing.T) {
	fibRedisRepo := SetupRedis()

	t.Run("Add", func(t *testing.T) {
		key := "5"
		value := "5"

		err := fibRedisRepo.Add(context.Background(), key, value)
		require.NoError(t, err)
		require.Nil(t, err)

		res, exist, err := fibRedisRepo.CheckFib(context.Background(), key)
		require.Equal(t, value, res)
		require.NoError(t, err)
		require.True(t, exist)
	})
}

func TestRedisRepo_CheckFib(t *testing.T) {
	fibRedisRepo := SetupRedis()

	t.Run("CheckFib", func(t *testing.T) {
		key := "5"
		value := "5"

		resBefore, exist, _ := fibRedisRepo.CheckFib(context.Background(), key)
		require.Equal(t, "", resBefore)
		require.False(t, exist)

		err := fibRedisRepo.Add(context.Background(), key, value)
		require.NoError(t, err)
		require.Nil(t, err)

		resAfter, exist, err := fibRedisRepo.CheckFib(context.Background(), key)
		require.Equal(t, value, resAfter)
		require.NoError(t, err)
		require.True(t, exist)
	})
}

// func TestRedisRepo_CheckFibAfter10Sec(t *testing.T) {
// 	fibRedisRepo := SetupRedis()

// 	t.Run("CheckFib", func(t *testing.T) {
// 		key := "5"
// 		value := "5"

// 		resBefore, exist, _ := fibRedisRepo.CheckFib(context.Background(), key)
// 		require.Equal(t, "", resBefore)
// 		require.False(t, exist)

// 		err := fibRedisRepo.Add(context.Background(), key, value)
// 		require.NoError(t, err)
// 		require.Nil(t, err)

// 		time.Sleep(6 * time.Second)

// 		resAfter, exist, _ := fibRedisRepo.CheckFib(context.Background(), key)
// 		require.Equal(t, "", resAfter)
// 		require.False(t, exist)
// 	})
// }