package service

import (
	"github.com/rht6226/clip-url/model"
)

type UrlService interface {
	Save(string) (string, error)
	Load(string) (*model.Url, error)
}
