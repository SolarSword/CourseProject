package cache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"

	cfg "course.project/authentication/config"
)

type Redis struct {
	redisClient *redis.Client
}

var Ca Redis

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

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		log.Fatalf("get error: %v when get data of %v from redis\n", err, key)
		return "", err
	}
	return val, nil
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := r.redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		log.Fatalf("get error: %v when set data of %v, %v from redis\n", err, key, value)
		return err
	}
	return nil
}

func (r *Redis) Del(ctx context.Context, key string) error {
	_, err := r.redisClient.Del(ctx, key).Result()
	if err != nil {
		log.Fatalf("get error: %v when delete key of %v from redis\n", err, key)
		return err
	}
	return nil
}
