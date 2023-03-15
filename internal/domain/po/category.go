package po

type Category struct {
	Id   int64  `gorm:"primaryKey;autoIncrement;comment:文章 ID;column:article_id"`
	Name string `gorm:"type:varchar(20);not null;comment:分类名称" json:"name"`
}

func (*Category) TableName() string {
	return "category"
}
