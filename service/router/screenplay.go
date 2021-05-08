package router

import (
	"backend/screenplay"
	"backend/source"
)

func registerScreenplayRouter() {
	svr := screenplay.GetService()
	router := e.Group("/screenplay")
	router.Use(source.AuthMiddleware())
	router.GET("/Info", svr.GetOneScreenplayInfo)
}
