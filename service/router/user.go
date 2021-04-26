package router

import "backend/user"

func registerUserRouter() {
	userRouter := e.Group("/user")
	userSvr := user.GetService()
	userRouter.PUT("/register", userSvr.RegisterUser)
	userRouter.GET("/login", userSvr.Login)
	userRouter.GET("/sendVerificationCode", userSvr.SendVerificationCode)
}
