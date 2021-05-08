package router

import (
	"backend/source"
	"backend/user"
)

func registerUserRouter() {
	svr := user.GetService()
	router := e.Group("/user")
	router.Use(source.AuthMiddleware())
	router.GET("/getUserInfo", svr.GetUserInfo)
	e.GET("/login", user.GetService().Login)
	e.PUT("/register", svr.RegisterUser)
	e.GET("/sendVerificationCode", svr.SendVerificationCode)
}
