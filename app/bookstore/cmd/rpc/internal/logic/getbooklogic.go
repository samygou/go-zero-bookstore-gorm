package logic

import (
	"context"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/svc"
	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"
	"go-zero-bookstore/common/logx"
)

type GetBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookLogic {
	return &GetBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取book
func (l *GetBookLogic) GetBook(in *pb.GetBookReq) (*pb.GetBookResp, error) {
	exist, err := l.svcCtx.Repo.ExistBookByID(l.ctx, in.Id)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if !exist {
		return nil, ErrBookNotExist
	}
	book, err := l.svcCtx.Repo.GetBook(l.ctx, in.Id)
	if err != nil {
		logx.Error(err)
		return nil, ErrInternalFault
	}

	return &pb.GetBookResp{
		Book: &pb.Book{
			Id:    book.ID,
			Name:  book.Name,
			Price: book.Price,
			Desc:  book.Desc,
		},
	}, nil
}
