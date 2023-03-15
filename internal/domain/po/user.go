package po

type User struct {
	Id           int64  `gorm:"primaryKey;autoIncrement;comment:用户ID;column:user_id"`
	NickName     string `gorm:"uniqueIndex;type:varchar(50);comment:用户名;column:nick_name"`
	Password     string `gorm:"type:varchar(100);comment:密码;column:password"`
	Salt         string `gorm:"type:varchar(100);comment:密码盐;column:salt"`
	Avatar       string `gorm:"type:varchar(1024);not null;comment:头像地址" json:"avatar"`
	Introduction string `gorm:"type:varchar(255);comment:个人简介" json:"introduction"`
	Email        string `gorm:"type:varchar(30);comment:邮箱" json:"email"`
}

func (*User) TableName() string {
	return "user"
}
