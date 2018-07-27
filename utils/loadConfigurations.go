package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	RedisConfig struct {
		RedisAddress         string `json:"redis_address"`
		RedisPoolMaxIdle     int    `json:"redis_pool_max_idle"`
		RedisPoolMaxActive   int    `json:"redis_pool_max_active"`
		RedisPoolIdleTimeout int    `json:"redis_pool_idle_timeout"`
		RedisPassword        string `json:"redis_password"`
	}
	KafkaConfig struct {
		KafkaBootStrapServer string `json:"bootstrap.servers"`
		KafkaGroupId         string `json:"group.id"`
		KafkaAutoOffSetReset string `json:"auto.offset.reset"`
		KafkaTopic           string `json:"kafka.topic"`
	}
}

var AllConfig = Config{}

func unmarshal(str string) (interface{}, error) {
	var iFace interface{}
	decoder := json.NewDecoder(strings.NewReader(str))
	decoder.UseNumber()
	if err := decoder.Decode(&iFace); err != nil {
		return nil, err
	}
	fmt.Println("This is Json Final ", iFace)
	return iFace, nil
}

func LoadConfigJson(fileName string) {
	confFile, err := os.Open(fileName)
	defer confFile.Close()
	if err != nil {
		panic("File read operation Failed")
	}
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&AllConfig)
}
func Get_R_Address() string {
	return AllConfig.RedisConfig.RedisAddress
}
func Get_R_Password() string {
	return AllConfig.RedisConfig.RedisPassword
}
func Get_R_PoolMaxActive() int {
	return AllConfig.RedisConfig.RedisPoolMaxActive
}
func Get_R_PoolMaxIdle() int {
	return AllConfig.RedisConfig.RedisPoolMaxIdle
}
func Get_R_PoolIdleTimeout() int {
	return AllConfig.RedisConfig.RedisPoolIdleTimeout
}

func Get_K_BootStrapServer() string {
	return AllConfig.KafkaConfig.KafkaBootStrapServer
}

func Get_K_GroupId() string {
	return AllConfig.KafkaConfig.KafkaGroupId
}

func Get_K_AutoOffSetReset() string {
	return AllConfig.KafkaConfig.KafkaAutoOffSetReset
}

func Get_K_Topic() string {
	return AllConfig.KafkaConfig.KafkaTopic
}
