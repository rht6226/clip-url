package service

import (
	"context"

	"github.com/rht6226/clip-url/model"
	"github.com/rht6226/clip-url/repository"
)

type urlService struct {
	repository repository.UrlRepository
}

func NewUrlService(repo repository.UrlRepository) *urlService {
	return &urlService{
		repository: repo,
	}
}

// save a new url into the db
func (svc *urlService) Save(ctx context.Context, url string) (encoded string, err error) {
	return "", nil
}

// load the url from the encoded short url 
func (svc *urlService) Load(ctx context.Context, shortLink string) (url *model.Url, err error) {
	return nil, nil
}