package svc

import (
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/config"
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/interfaces"
	repository2 "go-zero-bookstore/app/usercenter/cmd/rpc/internal/repository"
)

type ServiceContext struct {
	Config config.Config
	Repo   interfaces.AccountRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repository2.NewAccountRepo(),
	}
}
