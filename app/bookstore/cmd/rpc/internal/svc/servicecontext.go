package svc

import (
	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/config"
	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/interfaces"
	repository2 "go-zero-bookstore/app/bookstore/cmd/rpc/internal/repository"
)

type ServiceContext struct {
	Config config.Config
	Repo   interfaces.BookRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Repo:   repository2.NewBookRepo(),
	}
}
