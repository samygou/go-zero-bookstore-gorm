// Code generated by goctl. DO NOT EDIT.
// Source: bookstore.proto

package bookstore

import (
	"context"

	"go-zero-bookstore/app/bookstore/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Book            = pb.Book
	CreateBookReq   = pb.CreateBookReq
	CreateBookResp  = pb.CreateBookResp
	DeleteBookReq   = pb.DeleteBookReq
	DeleteBookResp  = pb.DeleteBookResp
	GetBookListReq  = pb.GetBookListReq
	GetBookListResp = pb.GetBookListResp
	GetBookReq      = pb.GetBookReq
	GetBookResp     = pb.GetBookResp
	UpdateBookReq   = pb.UpdateBookReq
	UpdateBookResp  = pb.UpdateBookResp

	Bookstore interface {
		// 添加book
		CreateBook(ctx context.Context, in *CreateBookReq, opts ...grpc.CallOption) (*CreateBookResp, error)
		// 更新book
		UpdateBook(ctx context.Context, in *UpdateBookReq, opts ...grpc.CallOption) (*UpdateBookResp, error)
		// 删除book
		DeleteBook(ctx context.Context, in *DeleteBookReq, opts ...grpc.CallOption) (*DeleteBookResp, error)
		// 获取book
		GetBook(ctx context.Context, in *GetBookReq, opts ...grpc.CallOption) (*GetBookResp, error)
		// 获取book list
		GetBookList(ctx context.Context, in *GetBookListReq, opts ...grpc.CallOption) (*GetBookListResp, error)
	}

	defaultBookstore struct {
		cli zrpc.Client
	}
)

func NewBookstore(cli zrpc.Client) Bookstore {
	return &defaultBookstore{
		cli: cli,
	}
}

// 添加book
func (m *defaultBookstore) CreateBook(ctx context.Context, in *CreateBookReq, opts ...grpc.CallOption) (*CreateBookResp, error) {
	client := pb.NewBookstoreClient(m.cli.Conn())
	return client.CreateBook(ctx, in, opts...)
}

// 更新book
func (m *defaultBookstore) UpdateBook(ctx context.Context, in *UpdateBookReq, opts ...grpc.CallOption) (*UpdateBookResp, error) {
	client := pb.NewBookstoreClient(m.cli.Conn())
	return client.UpdateBook(ctx, in, opts...)
}

// 删除book
func (m *defaultBookstore) DeleteBook(ctx context.Context, in *DeleteBookReq, opts ...grpc.CallOption) (*DeleteBookResp, error) {
	client := pb.NewBookstoreClient(m.cli.Conn())
	return client.DeleteBook(ctx, in, opts...)
}

// 获取book
func (m *defaultBookstore) GetBook(ctx context.Context, in *GetBookReq, opts ...grpc.CallOption) (*GetBookResp, error) {
	client := pb.NewBookstoreClient(m.cli.Conn())
	return client.GetBook(ctx, in, opts...)
}

// 获取book list
func (m *defaultBookstore) GetBookList(ctx context.Context, in *GetBookListReq, opts ...grpc.CallOption) (*GetBookListResp, error) {
	client := pb.NewBookstoreClient(m.cli.Conn())
	return client.GetBookList(ctx, in, opts...)
}