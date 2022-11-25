package config

import (
	"github.com/go-redis/redis"
	"log"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	result, err := rdb.Ping().Result()
	if err != nil {
		log.Print(err)
	}
	log.Println(result)
}

func GetRedisConnect() *redis.Client {
	return rdb
}
