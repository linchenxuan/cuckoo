package _interface

import (
	"cuckoo/internal/domain/application"
	"cuckoo/internal/domain/repository"
	"cuckoo/pkg/app"
	"cuckoo/pkg/cuckoo_error"
	"cuckoo/pkg/valobj"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleInterface struct {
	articleApp *application.ArticleApplication
}

func InitArticleInterface(engine *gin.Engine, articleRepo repository.IArticleRepository, tagRepo repository.ITagRepository, categoryRepo repository.ICategoryRepository) *ArticleInterface {
	articleInter := &ArticleInterface{
		articleApp: application.NewArticleApplication(articleRepo, tagRepo, categoryRepo),
	}

	lev := engine.Group("/v1/article")
	lev.GET("/:id", articleInter.GetArticle)
	lev.GET("/list", articleInter.GetArticleList)

	lev.PUT("", articleInter.CreateArticle)

	return articleInter
}

func (inter *ArticleInterface) GetArticle(ctx *gin.Context) {
	cuckoo := app.NewApplication(ctx)
	var idstr string
	param := map[string]*string{
		"id": &idstr,
	}
	cuckoo.InitRequestFromParam(param)

	id, err := strconv.Atoi(idstr)
	if err != nil {
		cuckoo.ErrorResponse(cuckoo_error.InvalidParams)
		return
	}
	articles, cerr := inter.articleApp.GetArticle(int64(id))

	cuckoo.SimpleResponse(cerr, articles)
}

func (inter *ArticleInterface) GetArticleList(ctx *gin.Context) {
	cuckoo := app.NewApplication(ctx)
	var req valobj.GetArticleListReq
	cuckoo.InitRequestFromBody(&req)

	inter.articleApp.GetArticleList(req)
}

func (inter *ArticleInterface) CreateArticle(ctx *gin.Context) {
	cuckoo := app.NewApplication(ctx)

	userInfo := cuckoo.GetUserInfo()
	if userInfo == nil {
		cuckoo.ErrorResponse(cuckoo_error.InternalError)
		return
	}

	var req valobj.CreateArticleReq
	if err := cuckoo.InitRequestFromBody(&req); err != nil {
		cuckoo.ErrorResponse(err)
		return
	}

	inter.articleApp.CreateArticle(userInfo.Id, req)
}
