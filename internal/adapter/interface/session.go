package _interface

import (
	"cuckoo/internal/domain/application"
	"cuckoo/internal/domain/repository"
	"cuckoo/internal/domain/service"
	"cuckoo/pkg/app"
	"cuckoo/pkg/valobj"
	"github.com/gin-gonic/gin"
)

type SessionInterface struct {
	sessionApp *application.SessionApplication
}

func InitSessionInterface(engine *gin.Engine, repo repository.IUserRepository) *SessionInterface {
	sessionInter := &SessionInterface{
		sessionApp: application.NewSessionApplication(service.SessionServiceOpt{
			Issuer: "",
			Secret: "",
			Expire: 0,
		}, repo),
	}

	lev := engine.Group("/v1/session")
	lev.PUT("", sessionInter.Login)

	return sessionInter
}

// 登录
func (inter *SessionInterface) Login(ctx *gin.Context) {
	cuckoo := app.NewApplication(ctx)
	var req valobj.LoginReq
	cuckoo.InitRequestFromBody(&req)

	inter.sessionApp.Login(req)
}
