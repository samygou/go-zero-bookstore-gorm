package logic

import (
	"context"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *usercenter.GetUserInfoReq) (*usercenter.GetUserInfoResp, error) {
	account, err := l.svcCtx.Repo.GetAccountInfo(l.ctx, in.Id)
	if err != nil {
		return nil, ErrAccountInternalFault
	}

	userInfo := &usercenter.User{
		Id:       account.ID,
		Mobile:   account.Mobile,
		Username: account.Username,
		Sex:      account.Sex,
		Avatar:   account.Avatar,
		Remark:   account.Remark,
	}

	return &usercenter.GetUserInfoResp{
		User: userInfo,
	}, nil
}
