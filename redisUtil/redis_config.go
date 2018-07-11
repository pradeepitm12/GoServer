package redisUtil

import (
	"encoding/json"
	"os"
)

type RedisConfig struct {
	RedisAddress         string `json:"redis_address"`
	RedisPoolMaxIdle     int    `json:"redis_pool_max_idle"`
	RedisPoolMaxActive   int    `json:"redis_pool_max_active"`
	RedisPoolIdleTimeout int    `json:"redis_pool_idle_timeout"`
	RedisPassword	string `json:"redis_password"`
}

var redisConf = RedisConfig{}

func LoadRedisConf(fileName string) (RedisConfig, error) {
	confFile, err := os.Open(fileName)
	defer confFile.Close()
	if err != nil {
		return redisConf, err
	}
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&redisConf)
	return redisConf, err
}

func GetRedisAddress() string {
	return redisConf.RedisAddress
}

func GetRedisMaxIdle() int {
	return redisConf.RedisPoolMaxIdle
}

func GetRedisMaxActive() int {
	return redisConf.RedisPoolMaxActive
}

func GetRedisIdleTimeout() int {
	return redisConf.RedisPoolIdleTimeout
}
func GetRedisPassword() string{
	return redisConf.RedisPassword
}
