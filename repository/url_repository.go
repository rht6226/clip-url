package repository

import (
	"github.com/rht6226/clip-url/model"
)

type UrlRepository interface {
	IsUsed(id uint64) bool
	Save(*model.Url) error
	Load(id uint64) (*model.Url, error)
	Close() error
}
