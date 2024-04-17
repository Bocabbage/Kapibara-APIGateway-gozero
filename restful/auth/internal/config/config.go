package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type MySQLConfig struct {
	MySQLConnStr  string
	MySQLRoleSalt int64
	// MySQLCacheConfig cache.CacheConf
}

type JWTConfig struct {
	JwtExpired   int64
	JwtSecretKey string
}

type CookieConfig struct {
	CookieServerDomain string
}

type CorsConfig struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
}

type Config struct {
	rest.RestConf
	MySQLConfig
	JWTConfig
	CookieConfig
	CorsConfig
}
