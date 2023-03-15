package po

import (
	"cuckoo/pkg/valobj"
)

type Article struct {
	Id          int64                `gorm:"primaryKey;autoIncrement;comment:文章 ID;column:article_id"`
	CategoryId  int64                `gorm:"type:bigint;not null;comment:分类 ID" json:"category_id"`
	UserId      int64                `gorm:"type:int;not null;comment:用户 ID" json:"user_id"`
	Title       string               `gorm:"type:varchar(100);not null;comment:文章标题" json:"title"`
	Description string               `gorm:"type:varchar(200);comment:文章描述" json:"desc"`
	Content     string               `gorm:"type:longtext;comment:文章内容" json:"content"`
	CoverImage  string               `gorm:"type:varchar(100);comment:封面图片地址" json:"img"`
	Type        valobj.ArticleType   `gorm:"type:tinyint;comment:类型(1-原创 2-转载 3-翻译)" json:"type"`
	Status      valobj.ArticleStatus `gorm:"type:tinyint;comment:状态(1-公开 2-私密)" json:"status"`
	OriginalUrl string               `gorm:"type:varchar(100);comment:源链接" json:"original_url"`
	IsTop       bool                 `gorm:"type:tinyint;not null;default:0;comment:是否置顶(0-否 1-是)" json:"is_top"`
	IsDelete    bool                 `gorm:"type:tinyint;not null;default:0;comment:是否放到回收站(0-否 1-是)" json:"is_delete"`
	Tags        []int64              `gorm:"type:int;comment:文章标签"`
}

func (*Article) TableName() string {
	return "article"
}
