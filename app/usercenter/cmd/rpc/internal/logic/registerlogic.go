package logic

import (
	"context"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/interfaces"
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/usercenter/cmd/rpc/usercenter"
	"go-zero-bookstore/common/logx"
	tool "go-zero-bookstore/common/tools"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(in *usercenter.RegisterReq) (*usercenter.RegisterResp, error) {
	if len(in.Username) == 0 {
		return nil, ErrUsernameIsEmpty
	}

	if len(in.Password) == 0 {
		return nil, ErrPasswordIncorrect
	}

	if in.Password != in.Password2 {
		return nil, ErrPasswordUnEqual
	}

	exist, err := l.svcCtx.Repo.ExistAccountByMobile(l.ctx, in.Mobile)
	if err != nil {
		return nil, ErrAccountInternalFault
	}
	if exist {
		return nil, ErrAccountAlreadyExists
	}

	accountId, err := l.svcCtx.Repo.CreateAccount(l.ctx, &interfaces.CreateAccountReq{
		Mobile:   in.Mobile,
		Username: in.Username,
		Password: tool.Md5ByString(in.Password),
		Sex:      0,
		Avatar:   "",
		Remark:   "",
	})
	if err != nil {
		return nil, ErrAccountInternalFault
	}

	//Generate token
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&usercenter.GenerateTokenReq{UserId: accountId})
	if err != nil {
		logx.Error(err)
		return nil, ErrAccountInternalFault
	}

	return &usercenter.RegisterResp{Token: tokenResp.Token}, nil
}
