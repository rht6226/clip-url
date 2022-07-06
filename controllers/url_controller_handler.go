package controllers

import (
	"github.com/gin-gonic/gin"
	service "github.com/rht6226/clip-url/service/url"
)

type urlControllerHandler struct {
	urlService service.UrlService
}

func NewUrlControllerHandler(r *gin.Engine, svc service.UrlService) {
	h := &urlControllerHandler{
		urlService: svc,
	}

	g := r.Group("")

	g.GET("/:shortLink/info", h.LoadUrlInfo)
	g.POST("/encode", h.SaveNewUrl)
	g.GET("/:shortLink", h.Redirect)
}

func (h *urlControllerHandler) LoadUrlInfo(c *gin.Context) {}

func (h *urlControllerHandler) SaveNewUrl(c *gin.Context) {}

func (h *urlControllerHandler) Redirect(c *gin.Context) {}
