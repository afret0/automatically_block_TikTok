package router

import "backend/user"

func registerUserRouter() {
	router := e.Group("/user")
	svr := user.GetService()
	router.PUT("/register", svr.RegisterUser)
	router.GET("/login", svr.Login)
	router.GET("/sendVerificationCode", svr.SendVerificationCode)
	router.GET("/getUserInfo", svr.GetUserInfo)
}
