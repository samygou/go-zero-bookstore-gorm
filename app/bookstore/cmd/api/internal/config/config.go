package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JWTAuth struct {
		AccessSecret  string
		AccessExpired int64
	}
	Bookstore zrpc.RpcClientConf
	Mysql     struct {
		DataSource      string
		MaxOpenConns    int
		MaxIdleConns    int
		MaxConnLifeTime int64
	}
}
