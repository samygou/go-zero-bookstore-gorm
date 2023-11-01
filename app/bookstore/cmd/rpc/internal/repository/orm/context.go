package orm

import "gorm.io/gorm"

type Context struct {
	DB *gorm.DB // 数据库engine
	tx *gorm.DB // 事务engine
}

func NewContext(db *gorm.DB) *Context {
	return &Context{DB: db}
}

func (c *Context) TX() *gorm.DB {
	return c.tx
}

func (c *Context) BeginTx() {
	if c.DB == nil {
		panic("DB init failed!!!")
	}

	c.tx = c.DB.Begin()
}

func (c *Context) Commit() error {
	if c.tx == nil {
		return nil
	}

	return c.tx.Commit().Error
}

func (c *Context) Rollback() {
	if c.tx == nil {
		return
	}

	c.tx.Rollback()
}
