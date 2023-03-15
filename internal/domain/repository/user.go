package repository

import (
	"cuckoo/internal/domain/po"
)

type IUserRepository interface {
	Get(id int64) (user *po.User, err error)
	FindByName(userName string) (user *po.User, err error)
	Save(user *po.User) error
	New(nickName string, password string, salt string) (user *po.User, err error)
}
