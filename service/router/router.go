package router

import (
	"backend/source"
	"github.com/gin-gonic/gin"
)

var e *gin.Engine

func init() {
	e = source.GetGinEngine()
}

func RegisterRouter() {
	registerUserRouter()
	registerTestRouter()
	registerScreenplayRouter()
}

func registerTestRouter() {
	e.GET("/test", func(c *gin.Context) {
		source.GetLogger().Infoln(source.Config.Get("mongo"))
		c.JSON(200, gin.H{"hello": "world"})
	})
}
