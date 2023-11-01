package repository

import (
	"context"

	"gorm.io/gorm"

	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/interfaces"
	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/repository/orm"
	"go-zero-bookstore/app/bookstore/cmd/rpc/internal/repository/tables"
	"go-zero-bookstore/common/logx"
)

type bookRepo struct {
	ctx *orm.Context
}

func NewBookRepo(db *gorm.DB) interfaces.BookRepo {
	return &bookRepo{
		ctx: orm.NewContext(db),
	}
}

func (repo *bookRepo) ExistBookByID(ctx context.Context, id int64) (bool, error) {
	type book struct {
		ID int64 `gorm:"id"`
	}

	var record book

	err := repo.ctx.DB.Model(&tables.Book{}).Select("id").Where("id = ?", id).Find(&record).Error
	if err != nil {
		return false, err
	}

	if record.ID == id {
		return true, nil
	}

	return false, nil
}

func (repo *bookRepo) CreateBook(ctx context.Context, req *interfaces.CreateBookReq) (int64, error) {
	book := tables.Book{
		Name:  req.Name,
		Price: req.Price,
		Desc:  req.Desc,
	}

	// 以下测试事务
	logx.Info(repo.ctx.TX())
	repo.ctx.BeginTx()
	defer func() {
		if r := recover(); r != nil {
			repo.ctx.Rollback()
		}
	}()
	err := repo.ctx.TX().Create(&book).Error
	if err != nil {
		logx.Error(err)
		return 0, err
	}
	logx.Info(repo.ctx.TX())

	return book.ID, repo.ctx.Commit()
}

func (repo *bookRepo) UpdateBook(ctx context.Context, req *interfaces.UpdateBookReq) (int64, error) {
	// 开启事务
	repo.ctx.BeginTx()
	defer func() {
		if r := recover(); r != nil {
			repo.ctx.Rollback()
		}
	}()

	err := repo.ctx.TX().Model(&tables.Book{}).Where("id = ?", req.ID).Updates(map[string]interface{}{
		"name":  req.Name,
		"price": req.Price,
		"desc":  req.Desc,
	}).Error
	if err != nil {
		logx.Error(err)
		return 0, err
	}

	return req.ID, repo.ctx.Commit()
}

func (repo *bookRepo) DeleteBook(ctx context.Context, id int64) (int64, error) {
	err := repo.ctx.DB.Where("id = ?", id).Delete(&tables.Book{}).Error
	if err != nil {
		logx.Error(err)
		return 0, err
	}
	return id, nil
}

func (repo *bookRepo) GetBook(ctx context.Context, id int64) (*interfaces.Book, error) {
	var book tables.Book

	err := repo.ctx.DB.Where("id = ?", id).Find(&book).Error
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &interfaces.Book{
		ID:       book.ID,
		Name:     book.Name,
		Price:    book.Price,
		Desc:     book.Desc,
		CreateAt: book.CreateAt,
		UpdateAt: book.UpdateAt,
	}, nil
}

func (repo *bookRepo) GetBooks(ctx context.Context, req *interfaces.GetBooksReq) ([]interfaces.Book, error) {
	var books []tables.Book

	session := repo.ctx.DB.Model(&tables.Book{})
	if req.Name != nil {
		session = session.Where("name = ?", *req.Name)
	}

	if req.Price != nil {
		session = session.Where("price = ?", *req.Price)
	}

	if req.OrderBy != nil {
		session = session.Order(*req.OrderBy)
	}

	offset := (req.Page - 1) * req.PageSize

	err := session.Offset(int(offset)).Limit(int(req.PageSize)).Find(&books).Error
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	var out []interfaces.Book

	for _, book := range books {
		out = append(out, interfaces.Book{
			ID:       book.ID,
			Name:     book.Name,
			Price:    book.Price,
			Desc:     book.Desc,
			CreateAt: book.CreateAt,
			UpdateAt: book.UpdateAt,
		})
	}

	return out, nil
}
