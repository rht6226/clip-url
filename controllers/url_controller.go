package controllers

import "github.com/gin-gonic/gin"

type UrlController interface {
	SaveNewUrl(*gin.Context)
	LoadUrlInfo(*gin.Context)
	Redirect(*gin.Context)
}