package service

import (
	"time"

	app_errors "github.com/rht6226/clip-url/errors"
	"github.com/rht6226/clip-url/model"
	"github.com/rht6226/clip-url/repository"
	service "github.com/rht6226/clip-url/service/encodec"
	"github.com/rht6226/clip-url/service/idgen"
)

type urlService struct {
	idgenService   idgen.IdGeneratorService
	encoDecService service.EncoDecService
	repository     repository.UrlRepository
}

// create a new url service with repo, idgen and encodec services
func NewUrlService(idgenService idgen.IdGeneratorService,
	encoDecService service.EncoDecService, repo repository.UrlRepository) *urlService {
	return &urlService{
		idgenService:   idgenService,
		encoDecService: encoDecService,
		repository:     repo,
	}
}

// save a new url into the db
func (svc *urlService) Save(url string) (encoded string, err error) {
	newId := svc.idgenService.GetId()

	for used := true; used; used = svc.repository.IsUsed(newId) {
		newId = svc.idgenService.GetId()
	}

	expires := time.Now().AddDate(1, 0, 0)

	urlObject := &model.Url{
		Id:      newId,
		URL:     url,
		Expires: expires.Format("2006-01-02 15:04:05.728046 +0300 EEST"),
		Visits:  0,
	}

	encoded = svc.encoDecService.Encode(urlObject.Id)

	err = svc.repository.Save(urlObject)

	if err != nil {
		return "", app_errors.NewInternal(err.Error())
	}

	return encoded, nil
}

// load the url from the encoded short url
func (svc *urlService) Load(shortLink string) (url *model.Url, err error) {
	decodedId, err := svc.encoDecService.Decode(shortLink)
	if err != nil {
		return nil, app_errors.NewBadRequest("Invalid URL Provided")
	}

	return svc.repository.Load(decodedId)
}
