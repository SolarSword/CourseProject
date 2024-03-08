package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"

	cfg "course.project/authentication/config"
)

type Redis struct {
	redisClient *redis.Client
}

func (r *Redis) InitCache() error {
	r.redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Cfg.Cache.Addr,
		Password: cfg.Cfg.Cache.Password,
		DB:       cfg.Cfg.Cache.DB,
	})
	ctx := context.Background()
	err := r.redisClient.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("get error: %v when connecting redis: %v \n", err, cfg.Cfg.Cache)
		return err
	}
	return nil
}

func (r *Redis) CloseCache() {
	err := r.redisClient.Close()
	if err != nil {
		panic(err)
	}
}
