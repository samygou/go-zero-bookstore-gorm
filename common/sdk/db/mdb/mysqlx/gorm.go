package mysqlx

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"go-zero-bookstore/common/logx"
	"go-zero-bookstore/common/typex"
)

var Sess *gorm.DB

func New(dataSource string, args map[string]interface{}) *gorm.DB {
	logx.Info("new mysql: ", dataSource)

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，即不会在表名后添加复数s
		},
	})
	typex.MustNil(err)

	sqlDB, err := db.DB()
	typex.MustNil(err)

	if value, ok := args["maxOpenConns"]; ok {
		switch value := value.(type) {
		case int:
			sqlDB.SetMaxOpenConns(value)
		default:
			panic("The maxOpenConns Must been an int type")
		}
	}

	if value, ok := args["maxIdleConns"]; ok {
		switch value := value.(type) {
		case int:
			sqlDB.SetMaxIdleConns(value)
		default:
			panic("The maxIdleConns Must been an int type")
		}
	}

	if value, ok := args["maxConnLifeTime"]; ok {
		switch value := value.(type) {
		case int64:
			sqlDB.SetConnMaxIdleTime(time.Duration(value) * time.Second)
		default:
			panic("The maxConnLifeTime Must been an int64 type")
		}
	}

	// 加入数据库的定时探测
	func() {
		period := 30 * time.Second
		timer := time.NewTimer(period)
		go func() {
			for {
				select {
				case <-timer.C:
					if err = sqlDB.Ping(); err != nil {
						logx.Warnf("db lose connect: %v", err)
					} else {
						logx.Info("db connected")
					}
					timer.Reset(period)
				}
			}
		}()
	}()

	return db
}
