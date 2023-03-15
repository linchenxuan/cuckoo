package service

import (
	"cuckoo/internal/domain/entity"
	"cuckoo/pkg/crypto"
	"github.com/google/uuid"
)

func encryptPassword(password string) (string, string) {
	salt := uuid.New().String()
	return crypto.MD5(crypto.MD5(password) + salt), salt
}

func checkPassword(user *entity.User, password string) bool {
	return user.Password == crypto.MD5(crypto.MD5(password)+user.Salt)
}
