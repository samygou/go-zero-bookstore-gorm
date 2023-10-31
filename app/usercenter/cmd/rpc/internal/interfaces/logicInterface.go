package interfaces

import (
	"context"
)

type AccountRepo interface {
	ExistAccountByMobile(ctx context.Context, mobile string) (bool, error)
	ExistAccountByID(ctx context.Context, id int64) (bool, error)
	GetAccountByMobile(ctx context.Context, mobile string) (*Account, error)
	CreateAccount(ctx context.Context, req *CreateAccountReq) (int64, error)
	UpdateAccount(ctx context.Context, req *UpdateAccountReq) (int64, error)
	GetAccountInfo(ctx context.Context, id int64) (*Account, error)
}
