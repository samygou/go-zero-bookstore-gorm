package logic

import (
	"context"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"
	"go-zero-bookstore/common/logx"
)

type DeleteBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBookLogic {
	return &DeleteBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 删除book
func (l *DeleteBookLogic) DeleteBook(in *pb.DeleteBookReq) (*pb.DeleteBookResp, error) {
	exist, err := l.svcCtx.Repo.ExistBookByID(l.ctx, in.Id)
	if err != nil {
		logx.Error(err)
		return nil, ErrInternalFault
	}
	if !exist {
		return nil, ErrBookNotExist
	}

	id, err := l.svcCtx.Repo.DeleteBook(l.ctx, in.Id)
	if err != nil {
		logx.Error(err)
		return nil, ErrInternalFault
	}

	return &pb.DeleteBookResp{
		Id: id,
	}, nil
}
