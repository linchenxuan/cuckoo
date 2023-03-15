package _interface

import (
	"cuckoo/internal/domain/application"
	"cuckoo/internal/domain/repository"
	app "cuckoo/pkg/app"
	"cuckoo/pkg/valobj"
	"github.com/gin-gonic/gin"
)

type UserInterface struct {
	userApp *application.UserApplication
}

func InitUserInterface(engine *gin.Engine, repo repository.IUserRepository) *UserInterface {
	userInter := &UserInterface{
		userApp: application.NewUserApplication(repo),
	}

	lev := engine.Group("/v1/user")
	lev.PUT("", userInter.Register)
	lev.POST("", userInter.UpdateUser)

	return userInter
}

func (inter *UserInterface) Register(ctx *gin.Context) {
	cuckoo := app.NewApplication(ctx)
	var req valobj.RegisterUserReq
	cuckoo.InitRequestFromBody(&req)

	cuckoo.SimpleResponse(inter.userApp.Register(req), nil)
}

func (inter *UserInterface) UpdateUser(ctx *gin.Context) {
	cuckoo := app.NewApplication(ctx)
	var req valobj.UpdateUserReq
	cuckoo.InitRequestFromBody(&req)

	cuckoo.SimpleResponse(inter.userApp.Update(req), nil)
}
