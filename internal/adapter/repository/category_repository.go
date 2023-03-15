package repository

import (
	"cuckoo/internal/domain/po"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(opt RepositoryOpt) *CategoryRepository {
	return &CategoryRepository{
		DB: newGormInstance(opt, &po.Category{}),
	}
}

func (repo *CategoryRepository) New(name string) (*po.Category, error) {
	category := po.Category{
		Name: name,
	}
	err := repo.DB.FirstOrCreate(&category, category).Error

	return &category, err
}
