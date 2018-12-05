package dao

import (
	"github.com/go-redis/redis"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/common/config"
)

func Redis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     config.APPConfig.Redis.Addr,
		Password: config.APPConfig.Redis.Password,
		DB:       config.APPConfig.Redis.DB,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		log.Fatal(log.Fields{
			"redis":  pong,
			"err":    err,
			"detail": config.APPConfig.Redis,
		})

	}
	return client
}
