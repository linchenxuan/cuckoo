package application

import (
	"cuckoo/internal/domain/entity"
	"cuckoo/internal/domain/repository"
	"cuckoo/internal/domain/service"
	"cuckoo/pkg/cuckoo_error"
	"cuckoo/pkg/valobj"
)

type ArticleApplication struct {
	articleSrv *service.ArticleService
}

func NewArticleApplication(articleRepo repository.IArticleRepository, tagRepo repository.ITagRepository, categoryRepo repository.ICategoryRepository) *ArticleApplication {
	return &ArticleApplication{articleSrv: service.NewArticleService(articleRepo, tagRepo, categoryRepo)}
}

func (app *ArticleApplication) GetArticle(id int64) (*entity.Article, cuckoo_error.IErrorStatus) {
	return app.articleSrv.GetArticle(id)
}

func (app *ArticleApplication) GetArticleList(req valobj.GetArticleListReq) ([]*entity.Article, int64, cuckoo_error.IErrorStatus) {
	return app.articleSrv.GetArticleList(req)
}

func (app *ArticleApplication) CreateArticle(uid int64, req valobj.CreateArticleReq) (*entity.Article, cuckoo_error.IErrorStatus) {
	return app.articleSrv.CreateArticle(uid, req)
}
