package repository

import (
	"cuckoo/internal/domain/po"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(opt RepositoryOpt) *UserRepository {
	return &UserRepository{
		DB: newGormInstance(opt, &po.User{}),
	}
}

func (repo *UserRepository) Get(id int64) (*po.User, error) {
	user := po.User{Id: id}
	e := repo.find(&user)
	return &user, e
}

func (repo *UserRepository) FindByName(name string) (*po.User, error) {
	user := po.User{NickName: name}
	e := repo.find(&user)
	return &user, e
}

func (repo *UserRepository) Save(user *po.User) error {
	return repo.DB.Where(user.Id).Updates(user).Error
}

func (repo *UserRepository) New(nickName string, password string, salt string) (*po.User, error) {
	user := po.User{
		NickName: nickName,
		Password: password,
		Salt:     salt,
	}
	e := repo.DB.Create(&user).Error
	return &user, e
}

func (repo *UserRepository) find(user *po.User) error {
	return repo.DB.Where(user).Last(user).Error
}
