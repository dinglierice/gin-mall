package cache

import (
	"strconv"

	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"

	"mall/conf"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client

// InitCache 在中间件中初始化redis链接  防止循环导包，所以放在这里
func InitCache() {
	Redis()
}

// Redis 在中间件中初始化redis链接
func Redis() {
	db, _ := strconv.ParseUint(conf.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		//暂时改动一下，redis连不上先不panic
		//panic(err)
	}
	RedisClient = client
}
