package svc

import (
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/config"
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/interfaces"
	repository2 "go-zero-bookstore/app/usercenter/cmd/rpc/internal/repository"
	"go-zero-bookstore/common/sdk/db/mdb/mysqlx"
)

type ServiceContext struct {
	Config config.Config
	Repo   interfaces.AccountRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo: repository2.NewAccountRepo(mysqlx.New(c.DB.DataSource, map[string]interface{}{
			"maxOpenConns":    c.Mysql.MaxOpenConns,
			"maxIdleConns":    c.Mysql.MaxIdleConns,
			"maxConnLifeTime": c.Mysql.MaxConnLifeTime,
		})),
	}
}
