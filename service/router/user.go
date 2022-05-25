package router

import (
	"github.com/gin-gonic/gin"
	"service/middleware"
	"service/user"
)

func registerUserRouter(e *gin.Engine) {
	svr := user.GetService()
	router := e.Group("/user")
	router.Use(middleware.AuthMiddleware())
	e.POST("/register", svr.RegisterUser)
	e.GET("/sendVerificationCode", svr.SendVerificationCode)
}
