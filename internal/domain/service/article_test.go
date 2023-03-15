package service

import (
	"cuckoo/internal/adapter/repository"
	"cuckoo/pkg/valobj"
	"testing"
)

func TestArticleService_CreateArticle(t *testing.T) {
	repoOpt := repository.RepositoryOpt{
		User:     "lcx",
		Password: "123456",
		Host:     "192.168.163.128",
		Port:     3306,
		DBName:   "cuckoo",
	}
	articleRepo := repository.NewArticleRepository(repoOpt)
	tagRepo := repository.NewTagRepository(repoOpt)
	categoryRepo := repository.NewCategoryRepository(repoOpt)
	articleService := NewArticleService(articleRepo, tagRepo, categoryRepo)
	articleService.CreateArticle(1, valobj.CreateArticleReq{
		Title:    "test",
		Content:  "这是一篇测试文章",
		Type:     valobj.Original,
		Status:   valobj.PUBLIC,
		IsTop:    false,
		Category: "test",
	})
	articleService.GetArticleList(valobj.GetArticleListReq{
		PageSize: 10,
		PageNum:  1,
	})
}
