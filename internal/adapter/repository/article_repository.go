package repository

import (
	"cuckoo/internal/domain/po"
	"gorm.io/gorm"
)

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(opt RepositoryOpt) *ArticleRepository {
	return &ArticleRepository{
		DB: newGormInstance(opt, &po.Article{}),
	}
}

func (repo *ArticleRepository) Get(id int64) (*po.Article, error) {
	article := po.Article{Id: id}
	e := repo.find(&article)
	return &article, e
}

func (repo *ArticleRepository) FindsByPage(pageNum, pageSize, category, tag int) ([]*po.Article, int64, error) {
	articleList := make([]*po.Article, 0)
	var total int64

	db := repo.DB.Table("article").
		Select("id, title, content, img, type, is_top, created_at, category_id").
		Where("is_delete = 0 AND status = 1")
	if category != 0 {
		db = db.Where("category_id", category)
	}
	if tag != 0 {
		db = db.Where("id IN (SELECT article_id FROM article_tag WHERE tag_id = ?)", tag)
	}

	db.Count(&total)
	db.Preload("Tags").
		Preload("Category").
		Order("is_top DESC, id DESC").
		Limit(pageSize).Offset(pageSize * (pageNum - 1)).
		Find(&articleList)

	return articleList, total, nil
}

func (repo *ArticleRepository) find(user *po.Article) error {
	return repo.DB.Where(user).Last(user).Error
}

func (repo *ArticleRepository) New(article *po.Article) error {
	return repo.DB.Create(article).Error
}
