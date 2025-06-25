package interfaces

import "time"

type CacheProvider interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (string, error)
	Incr(key string) (int64, error)
	Expire(key string, ttl time.Duration) error
	TTL(key string) (time.Duration, error)
	Del(key string) error
}
