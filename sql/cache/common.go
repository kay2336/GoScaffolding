package cache

import (
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

// Redis
// 初始化redis链接
//func Redis() {
//	db, _ := strconv.ParseUint(conf.RedisDbName, 10, 64)
//	client := redis.NewClient(&redis.Options{
//		Addr:     conf.RedisAddr,
//		Password: conf.RedisPw,
//		DB:       int(db),
//	})
//	_, err := client.Ping().Result()
//	if err != nil {
//		logging.Info(err)
//		panic(err)
//	}
//	RedisClient = client
//}
