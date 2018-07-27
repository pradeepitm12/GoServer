package utils

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	redisPool *redis.Pool
)

func NewPool() *redis.Pool {
	config := AllConfig.RedisConfig
	return &redis.Pool{
		MaxIdle:     config.RedisPoolMaxIdle,
		MaxActive:   config.RedisPoolMaxActive,
		IdleTimeout: time.Duration(config.RedisPoolIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			//return redis.Dial("tcp", config.RedisAddress) },
			c, err := redis.Dial("tcp", config.RedisAddress)
			if err != nil {
				return nil, err
			}
			c.Do("AUTH", config.RedisPassword)
			return c, err
		},
	}
}

func InitRedisPool() {
	config := AllConfig.RedisConfig
	log.Println("Redis *********")
	log.Printf("Init Redis Connection pool with params: RedisAddress = %s, MaxIdle = %d, MaxActive = %d, IdleTimeout = %d(s)",
		config.RedisAddress, config.RedisPoolMaxIdle, config.RedisPoolMaxActive, config.RedisPoolIdleTimeout)
	redisPool = NewPool()
}
