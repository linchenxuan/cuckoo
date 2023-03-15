package application

import (
	"cuckoo/internal/domain/repository"
	"cuckoo/internal/domain/service"
	"cuckoo/pkg/cuckoo_error"
	"cuckoo/pkg/valobj"
)

type UserApplication struct {
	userSrv *service.UserService
}

func NewUserApplication(repo repository.IUserRepository) *UserApplication {
	return &UserApplication{userSrv: service.NewUserService(repo)}
}

func (app *UserApplication) Register(req valobj.RegisterUserReq) cuckoo_error.IErrorStatus {
	return app.userSrv.Register(req)
}

func (app *UserApplication) Update(req valobj.UpdateUserReq) cuckoo_error.IErrorStatus {
	return app.userSrv.UpdateUser(req)
}
