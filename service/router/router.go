package router

import (
	"github.com/gin-gonic/gin"
	"service/source"
)

func RegisterRouter(e *gin.Engine) {
	//registerUserRouter(e)
	registerTestRouter(e)
}

func registerTestRouter(e *gin.Engine) {
	e.GET("/test", func(c *gin.Context) {
		source.GetLogger().Infoln(source.Config.Get("mongo"))
		c.JSON(200, gin.H{"hello": "world"})
	})
}
