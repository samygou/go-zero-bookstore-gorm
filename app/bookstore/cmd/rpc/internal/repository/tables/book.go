package tables

type Book struct {
	ID       int64  `gorm:"column:id;int unsigned;primaryKey;autoIncrement;comment('主键id')"`
	CreateAt int64  `gorm:"column:create_at;bigint unsigned;notnull;autoCreateTime;comment('创建时间')"`
	UpdateAt int64  `gorm:"column:update_at;bigint unsigned;notnull;autoUpdateTime;comment('更新时间')"`
	Name     string `gorm:"column:name;varchar(32);notnull;default '';comment('书籍名称')"`
	Price    int64  `gorm:"column:price;bigint unsigned;notnull;default 0;comment('价格')"`
	Desc     string `gorm:"column:desc;varchar(255);notnull;default '';comment('备注')"`
}

func (Book) TableName() string {
	return "books"
}
