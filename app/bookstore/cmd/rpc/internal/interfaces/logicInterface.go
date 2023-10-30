package interfaces

import (
	"context"
	"gorm.io/gorm"
)

type BookRepo interface {
	ExistBookByID(ctx context.Context, sess *gorm.DB, id int64) (bool, error)
	CreateBook(ctx context.Context, sess *gorm.DB, req *CreateBookReq) (int64, error)
	UpdateBook(ctx context.Context, sess *gorm.DB, book *UpdateBookReq) (int64, error)
	DeleteBook(ctx context.Context, sess *gorm.DB, id int64) (int64, error)
	GetBook(ctx context.Context, sess *gorm.DB, id int64) (*Book, error)
	GetBooks(ctx context.Context, sess *gorm.DB, req *GetBooksReq) ([]Book, error)
}
