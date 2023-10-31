package interfaces

import (
	"context"
)

type BookRepo interface {
	ExistBookByID(ctx context.Context, id int64) (bool, error)
	CreateBook(ctx context.Context, req *CreateBookReq) (int64, error)
	UpdateBook(ctx context.Context, book *UpdateBookReq) (int64, error)
	DeleteBook(ctx context.Context, id int64) (int64, error)
	GetBook(ctx context.Context, id int64) (*Book, error)
	GetBooks(ctx context.Context, req *GetBooksReq) ([]Book, error)
}
