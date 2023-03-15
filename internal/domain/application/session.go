package application

import (
	"cuckoo/internal/domain/repository"
	"cuckoo/internal/domain/service"
	"cuckoo/pkg/cuckoo_error"
	"cuckoo/pkg/valobj"
)

type SessionApplication struct {
	UserSrv    *service.UserService
	SessionSrv *service.SessionService
}

func NewSessionApplication(opt service.SessionServiceOpt, repo repository.IUserRepository) *SessionApplication {
	return &SessionApplication{
		UserSrv:    service.NewUserService(repo),
		SessionSrv: service.NewSessionService(opt),
	}
}

func (app *SessionApplication) Login(req valobj.LoginReq) (string, cuckoo_error.IErrorStatus) {
	info, err := app.UserSrv.GetUserInfo(req.NickName)
	if err != nil {
		return "", cuckoo_error.InternalError
	}

	return app.SessionSrv.Login(req, info)
}
