package repository

import "cuckoo/internal/domain/po"

type IArticleRepository interface {
	New(*po.Article) error
	Get(id int64) (*po.Article, error)
	FindsByPage(pageNum, pageSize, category, tag int) ([]*po.Article, int64, error)
}
