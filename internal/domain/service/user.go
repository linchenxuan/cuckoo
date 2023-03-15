package service

import (
	"cuckoo/internal/domain/entity"
	"cuckoo/internal/domain/repository"
	"cuckoo/pkg/cuckoo_error"
	"cuckoo/pkg/valobj"
)

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (srv *UserService) Register(req valobj.RegisterUserReq) cuckoo_error.IErrorStatus {
	// TODO 并发问题
	if srv.userExited(req.Nickname) {
		return cuckoo_error.UserAlreadyExist
	}

	password, salt := encryptPassword(req.Password)

	_, err := srv.userRepo.New(req.Nickname, password, salt)
	if err != nil {
		return cuckoo_error.InternalError
	}

	return nil
}

func (srv *UserService) GetUserInfo(nickName string) (*entity.User, error) {
	pUser, err := srv.userRepo.FindByName(nickName)
	if err != nil {
		return nil, err
	}
	return &entity.User{User: *pUser}, nil
}

func (srv *UserService) UpdateUser(req valobj.UpdateUserReq) cuckoo_error.IErrorStatus {
	if srv.userExited(req.Nickname) {
		return cuckoo_error.UserAlreadyExist
	}

	user, err := srv.userRepo.Get(req.ID)
	if err != nil {
		return cuckoo_error.InternalError
	}
	if user == nil {
		return cuckoo_error.UserNotExist
	}

	user.NickName = req.Nickname
	user.Avatar = req.Avatar
	user.Email = req.Email
	user.Introduction = req.Introduction

	if srv.userRepo.Save(user) != nil {
		return cuckoo_error.InternalError
	}

	return nil
}

func (srv *UserService) userExited(nickName string) bool {
	user, _ := srv.userRepo.FindByName(nickName)
	return user != nil && user.Id > 0
}
