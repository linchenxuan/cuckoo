package service

import (
	"cuckoo/internal/domain/entity"
	"cuckoo/internal/utils"
	"cuckoo/pkg/cuckoo_error"
	"cuckoo/pkg/valobj"
	"time"
)

type SessionServiceOpt struct {
	Issuer string
	Secret string
	Expire time.Duration
}

type SessionService struct {
	opt SessionServiceOpt
}

func NewSessionService(opt SessionServiceOpt) *SessionService {
	return &SessionService{
		opt: opt,
	}
}

func (srv *SessionService) Login(req valobj.LoginReq, user *entity.User) (string, cuckoo_error.IErrorStatus) {
	if !checkPassword(user, req.Password) {
		return "", cuckoo_error.UserAuthFailed
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", cuckoo_error.InternalError
	}

	// TODO 更新登录时间、ip

	return token, nil
}
