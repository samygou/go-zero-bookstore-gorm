package tables

type Account struct {
	ID         int64  `gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement;comment:('主键id, 自增')"`
	CreateTime int64  `gorm:"column:create_time;type:bigint unsigned;autoCreateTime;comment('创建时间')"`
	UpdateTime int64  `gorm:"column:update_time;type:bigint unsigned;autoUpdateTime;comment('更新时间')"`
	DelStatus  int64  `gorm:"column:del_status;type:tinyint(1) unsigned;notnull;comment('删除状态')"`
	Mobile     string `gorm:"column:mobile;varchar(11);notnull;default '';comment('电话号码')"`
	Username   string `gorm:"column:username;varchar(32);notnull;default '';comment('用户名')"`
	Password   string `gorm:"column:password;varchar(255);notnull;comment('密码')"`
	Sex        int64  `gorm:"column:sex;tinyint(1) unsigned;notnull;default 0;comment('性别, 0未知, 1男性, 2女性')"`
	Avatar     string `gorm:"column:avatar;varchar(255);notnull;default '';comment('头像')"`
	Remark     string `gorm:"column:remark;varchar(255);notnull;default '';comment('备注')"`
}

func (Account) TableName() string {
	return "accounts"
}
