package app

import (
	ctx "context"
	"fmt"
	cacheifaces "pnBot/internal/cache/interfaces"
	redisprov "pnBot/internal/cache/redis"
	loaders "pnBot/internal/config/loaders"
	"pnBot/internal/config/models"
	loggerifaces "pnBot/internal/logger/interfaces"

	"github.com/redis/go-redis/v9"
)

func createRedisClient(
	context ctx.Context,
	cacheConfig *models.Cache,
	logger loggerifaces.Logger,
) cacheifaces.CacheProvider {
	cacheHost, cachePort, cacheUsername, cachePassword := loaders.LoadCacheConfig(*cacheConfig)

	cacheAddress := fmt.Sprintf("%s:%s", cacheHost, cachePort)

	redisClientOptions := redis.Options{
		Addr:     cacheAddress,
		Username: cacheUsername,
		Password: cachePassword,
	}

	redisClient := redis.NewClient(&redisClientOptions)
	if err := redisClient.Ping(context).Err(); err != nil {
		logger.Fatalf("не удалось подключиться к Redis: %v", err)
	}

	logger.Infof("Redis-клиент запущен на: %s", cacheAddress)

	redisCacheProvider := redisprov.NewRedisCacheProvider(redisClient, context)
	go func() {
		<-context.Done()
		if err := redisCacheProvider.GracefulShutdown(); err != nil {
			logger.Errorf("Ошибка при завершении Redis-клиента: %v", err)
		} else {
			logger.Info("Redis-клиент успешно завершен")
		}
	}()

	return redisCacheProvider
}
