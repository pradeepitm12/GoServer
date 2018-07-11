package redisUtil

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

func NewPool(config *RedisConfig) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     config.RedisPoolMaxIdle,
		MaxActive:   config.RedisPoolMaxActive,
		IdleTimeout: time.Duration(config.RedisPoolIdleTimeout) * time.Second,
		Dial:        func() (redis.Conn, error) {
			 //return redis.Dial("tcp", config.RedisAddress) },
			 c,err:=redis.Dial("tcp", config.RedisAddress)
			 if err != nil {
					return nil, err
				}
			c.Do("AUTH",config.RedisPassword)
		return c,err},
	}
}

func InitRedisPool(config *RedisConfig) {
	log.Printf("Init Redis Connection pool with params: RedisAddress = %s, MaxIdle = %d, MaxActive = %d, IdleTimeout = %d(s)",
		config.RedisAddress, config.RedisPoolMaxIdle, config.RedisPoolMaxActive, config.RedisPoolIdleTimeout)
	pool = NewPool(config)
}
