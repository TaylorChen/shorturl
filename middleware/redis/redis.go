package redis

import (
	"github.com/go-redis/redis"
	"log"
	"shorturl/conf"
)

var (
	redisdb *redis.Client
)

func Redis() *redis.Client {
	if redisdb == nil {
		newRedis, err := NewRedis()
		if err != nil {
			log.Println(err)
		}
		redisdb = newRedis
	}
	return redisdb
}

func NewRedis() (*redis.Client, error) {
	redisdb = redis.NewClient(&redis.Options{
		Addr:     conf.Conf.Redis.Server,
		Password: conf.Conf.Redis.Pwd, // no password set
		DB:       0,                   // use default DB
	})
	//defer Redis.Close()
	_, err := redisdb.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisdb, nil
}
