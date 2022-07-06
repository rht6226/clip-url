package service

import (
	"context"

	"github.com/rht6226/clip-url/model"
)

type UrlService interface {
	Save(context.Context, *model.Url) (string, error)
	Load(context.Context, string) (*model.Url, error)
}
