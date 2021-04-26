package router

import (
	"backend/utils"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() {
	E := utils.GetGinEngine()
	res := utils.GetResTemplateManager()

	userRouter := E.Group("/user")
	userRouter.PUT("/update", func(ctx *gin.Context) {
		//_ = user.GetManager().RegisterUser(ctx, "test", "18435155427")
		ctx.JSON(200, res.NewSucceedRes())

	})
	E.GET("/test", func(c *gin.Context) {
		utils.Logger.Info(utils.Config.Get("mongo"))
		c.JSON(200, gin.H{"hello": "world"})
	})
}
