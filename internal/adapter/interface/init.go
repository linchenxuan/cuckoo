package _interface

import (
	"cuckoo/internal/adapter/interface/middleware"
	"cuckoo/internal/adapter/repository"
	"cuckoo/internal/conf"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()

	engine.Use(middleware.Logger())

	repoOpt := conf.GetRepositoryOpt()
	userRepo := repository.NewUserRepository(repoOpt)
	articleRepo := repository.NewArticleRepository(repoOpt)
	tagRepo := repository.NewTagRepository(repoOpt)
	categoryRepo := repository.NewCategoryRepository(repoOpt)

	InitUserInterface(engine, userRepo)
	InitSessionInterface(engine, userRepo)
	InitArticleInterface(engine, articleRepo, tagRepo, categoryRepo)

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
