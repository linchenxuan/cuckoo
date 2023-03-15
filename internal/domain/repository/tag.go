package repository

import "cuckoo/internal/domain/po"

type ITagRepository interface {
	New(name string) (*po.Tag, error)
}
