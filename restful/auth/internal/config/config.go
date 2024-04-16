package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type MySQLConfig struct {
	MySQLConnStr     string
	MySQLPwdSalt     int64
	MySQLCacheConfig cache.CacheConf
}

type JWTConfig struct {
	JwtExpired   int64
	JwtSecretKey string
}

type ServerConfig struct {
	ServerAddr   string
	ServerPort   string
	ServerDomain string
	CertFile     string
	KeyFile      string
}

type Config struct {
	rest.RestConf
	MySQLConfig
	JWTConfig
	ServerConfig
}
