package cache

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

type Redis struct {
	Rdb *redis.Client
}

func NewRedisCLi() *Redis {
	return &Redis{Rdb: redis.NewClient(&redis.Options{Addr: "host.docker.internal:6379"})} //host.docker.internal
}

func (r *Redis) Set(pic []byte, key string) error {
	err := r.Rdb.Set(context.Background(), key, pic, 5*time.Minute).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(key string) ([]byte, error) {

	str, err := r.Rdb.Get(context.Background(), key).Result()
	if err != nil {
		return []byte(str), err
	}

	return []byte(str), nil
}
