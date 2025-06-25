package loaders

import (
	conf "pnBot/internal/config/models"
)

func LoadCacheConfig(cacheConfig conf.Cache) (string, string, string, string) {
	host := cacheConfig.Host
	port := cacheConfig.Port
	username := cacheConfig.Username
	password := cacheConfig.Password

	return host, port, username, password
}
