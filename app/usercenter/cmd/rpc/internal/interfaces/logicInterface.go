package interfaces

import (
	"context"
	"gorm.io/gorm"
)

type AccountRepo interface {
	ExistAccountByMobile(ctx context.Context, sess *gorm.DB, mobile string) (bool, error)
	ExistAccountByID(ctx context.Context, sess *gorm.DB, id int64) (bool, error)
	GetAccountByMobile(ctx context.Context, sess *gorm.DB, mobile string) (*Account, error)
	CreateAccount(ctx context.Context, sess *gorm.DB, req *CreateAccountReq) (int64, error)
	UpdateAccount(ctx context.Context, sess *gorm.DB, req *UpdateAccountReq) (int64, error)
	GetAccountInfo(ctx context.Context, sess *gorm.DB, id int64) (*Account, error)
}
