package rediscache

import (
	ctx "context"
	"errors"
	"time"

	cacheerrors "pnBot/internal/cache/errors"

	"github.com/redis/go-redis/v9"
)

type RedisCacheProvider struct {
	client  *redis.Client
	context ctx.Context
}

func NewRedisCacheProvider(redisClient *redis.Client, context ctx.Context) *RedisCacheProvider {
	return &RedisCacheProvider{
		client:  redisClient,
		context: context,
	}
}

func (rс *RedisCacheProvider) Set(key string, value interface{}, ttl time.Duration) error {
	return rс.client.Set(rс.context, key, value, ttl).Err()
}

func (rс *RedisCacheProvider) Get(key string) (string, error) {
	value, err := rс.client.Get(rс.context, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", cacheerrors.ErrNilVal
	}
	if err != nil {
		return "", err
	}
	return value, nil
}

func (rс *RedisCacheProvider) Incr(key string) (int64, error) {
	return rс.client.Incr(rс.context, key).Result()
}

func (rс *RedisCacheProvider) Expire(key string, ttl time.Duration) error {
	return rс.client.Expire(rс.context, key, ttl).Err()
}

func (rс *RedisCacheProvider) TTL(key string) (time.Duration, error) {
	return rс.client.TTL(rс.context, key).Result()
}

func (rс *RedisCacheProvider) Del(key string) error {
	return rс.client.Del(rс.context, key).Err()
}

func (rc *RedisCacheProvider) GracefulShutdown() error {
	return rc.client.Close()
}
