package router

import (
	"backend/script"
	"backend/source"
)

func registerScriptRouter() {
	svr := script.GetService()
	router := e.Group("/script")
	router.Use(source.AuthMiddleware())
	router.GET("/scriptInfo", svr.GetOneScriptInfo)

}
