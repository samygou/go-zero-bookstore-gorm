package repository

import (
	"context"

	"gorm.io/gorm"

	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/interfaces"
	"go-zero-bookstore/app/usercenter/cmd/rpc/internal/repository/tables"
	"go-zero-bookstore/common/logx"
)

type (
	accountRepo struct{}
)

func NewAccountRepo() interfaces.AccountRepo {
	return &accountRepo{}
}

func (repo *accountRepo) ExistAccountByMobile(ctx context.Context, sess *gorm.DB, mobile string) (bool, error) {
	type account struct {
		ID int64 `gorm:"id"`
	}

	var result account

	err := sess.Model(&tables.Account{}).Where("mobile = ?", mobile).Find(&result).Error
	if err != nil {
		logx.Error(err)
		return false, err
	}

	if result.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *accountRepo) ExistAccountByID(ctx context.Context, sess *gorm.DB, id int64) (bool, error) {
	type account struct {
		ID int64 `gorm:"id"`
	}

	var result account

	err := sess.Model(&tables.Account{}).Where("id = ?", id).Find(&result).Error
	if err != nil {
		logx.Error(err)
		return false, err
	}

	if result.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (repo *accountRepo) GetAccountByMobile(ctx context.Context, sess *gorm.DB, mobile string) (*interfaces.Account, error) {
	var account tables.Account

	err := sess.Where("mobile = ?", mobile).Find(&account).Error
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &interfaces.Account{
		ID:         account.ID,
		CreateTime: account.CreateTime,
		UpdateTime: account.UpdateTime,
		Mobile:     account.Mobile,
		Password:   account.Password,
		Sex:        account.Sex,
		Avatar:     account.Avatar,
		Remark:     account.Remark,
	}, nil
}

func (repo *accountRepo) CreateAccount(ctx context.Context, sess *gorm.DB, req *interfaces.CreateAccountReq) (int64, error) {
	account := tables.Account{
		DelStatus: 0,
		Mobile:    req.Mobile,
		Username:  req.Username,
		Password:  req.Password,
		Sex:       req.Sex,
		Avatar:    req.Avatar,
		Remark:    req.Remark,
	}
	err := sess.Create(&account).Error
	if err != nil {
		logx.Error(err)
		return 0, err
	}

	return account.ID, nil
}

func (repo *accountRepo) UpdateAccount(ctx context.Context, sess *gorm.DB, req *interfaces.UpdateAccountReq) (int64, error) {
	err := sess.Model(&tables.Account{}).Where("id = ?", req.ID).Updates(map[string]interface{}{
		"mobile":   req.Mobile,
		"username": req.Username,
		"sex":      req.Sex,
		"avatar":   req.Avatar,
		"remark":   req.Remark,
	}).Error
	if err != nil {
		logx.Error(err)
		return 0, err
	}

	return req.ID, nil
}

func (repo *accountRepo) GetAccountInfo(ctx context.Context, sess *gorm.DB, id int64) (*interfaces.Account, error) {
	var account tables.Account

	err := sess.Where("id = ?", id).Find(&account).Error
	if err != nil {
		logx.Error(err)
		return nil, err
	}

	return &interfaces.Account{
		ID:         account.ID,
		CreateTime: account.CreateTime,
		UpdateTime: account.UpdateTime,
		Mobile:     account.Mobile,
		Username:   account.Username,
		Sex:        account.Sex,
		Avatar:     account.Avatar,
		Remark:     account.Remark,
	}, nil
}
