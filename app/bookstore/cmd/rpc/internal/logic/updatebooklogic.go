package logic

import (
	"context"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/interfaces"
	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"
	"go-zero-bookstore/common/logx"
)

type UpdateBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateBookLogic {
	return &UpdateBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 更新book
func (l *UpdateBookLogic) UpdateBook(in *pb.UpdateBookReq) (*pb.UpdateBookResp, error) {
	if len(in.Name) == 0 {
		return nil, ErrBookNameIsNull
	}

	if in.Price <= 0 {
		return nil, ErrBookPriceIsIncorrect
	}

	exist, err := l.svcCtx.Repo.ExistBookByID(l.ctx, in.Id)
	if err != nil {
		return nil, ErrInternalFault
	}
	if !exist {
		return nil, ErrBookNotExist
	}

	id, err := l.svcCtx.Repo.UpdateBook(l.ctx, &interfaces.UpdateBookReq{
		ID:    in.Id,
		Name:  in.Name,
		Price: in.Price,
		Desc:  in.Desc,
	})
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &pb.UpdateBookResp{
		Id: id,
	}, nil
}
