package redis

import (
	// "time"

	"github.com/go-redis/redis/v8"
	"github.com/restlesswhy/grpc/grpc-rest-fibonacci-sequence/config"
)

// Returns new redis client
func NewRedisClient(cfg *config.Config) *redis.Ring {
	// client := redis.NewClient(&redis.Options{
	// 	Addr:         cfg.Redis.RedisAddr,
	// 	Password:     cfg.Redis.Password, // no password set
	// 	DB:           cfg.Redis.DB,       // use default DB
	// })
	
	ring := redis.NewRing(&redis.RingOptions{
        Addrs: map[string]string{
            "server1": cfg.Redis.RedisAddr,
        },
    })

	return ring
}
