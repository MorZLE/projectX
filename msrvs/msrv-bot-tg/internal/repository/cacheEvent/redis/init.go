package redis

import (
	"github.com/redis/go-redis/v9"
	"log/slog"
	"projectX/msrvs/msrv-bot-tg/config"
	"projectX/msrvs/msrv-bot-tg/internal/repository/cacheEvent"
)

func InitRedis(cnf config.Redis) cacheEvent.IStackEvent {
	opt, err := redis.ParseURL(cnf.Dsn)
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)
	slog.Info("Redis connected")
	return &RedisDB{db: client}
}

type RedisDB struct {
	db *redis.Client
}
