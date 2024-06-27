package core

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gvb_server/global"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConf := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Addr(),
		Password: redisConf.Password, // no password set
		DB:       db,                 // use default DB
		PoolSize: redisConf.PoolSize, // 连接池大小
	})
	context.Background()
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		logrus.Error("fail to connect to redis %s", redisConf.Addr())
		return nil
	}
	return rdb
}
