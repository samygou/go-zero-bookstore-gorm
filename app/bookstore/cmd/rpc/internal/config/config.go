package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	JWTAuth struct {
		AccessSecret  string
		AccessExpired int64
	}
	DB struct {
		DataSource string
	}
	Cache            cache.CacheConf
	DefaultPageParam struct {
		Page     int64
		PageSize int64
	}
	Mysql struct {
		DataSource      string
		MaxOpenConns    int
		MaxIdleConns    int
		MaxConnLifeTime int64
	}
}
