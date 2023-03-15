package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type RepositoryOpt struct {
	User     string
	Password string
	Host     string
	Port     int64
	DBName   string
}

func newGormInstance(opt RepositoryOpt, entitys ...interface{}) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", opt.User, opt.Password, opt.Host, opt.Port, opt.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	for _, entity := range entitys {
		if db.AutoMigrate(entity) != nil {
			panic(err)
		}
	}

	return db
}
