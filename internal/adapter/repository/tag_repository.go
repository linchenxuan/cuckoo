package repository

import (
	"cuckoo/internal/domain/po"
	"gorm.io/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(opt RepositoryOpt) *TagRepository {
	return &TagRepository{
		DB: newGormInstance(opt, &po.Tag{}),
	}
}

func (repo *TagRepository) New(name string) (*po.Tag, error) {
	tag := po.Tag{
		Name: name,
	}
	err := repo.DB.FirstOrCreate(&tag, tag).Error

	return &tag, err
}
