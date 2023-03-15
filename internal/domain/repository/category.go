package repository

import "cuckoo/internal/domain/po"

type ICategoryRepository interface {
	New(category string) (*po.Category, error)
}
