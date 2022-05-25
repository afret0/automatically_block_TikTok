package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"service/middleware"
	"service/router"
	"service/source"
	"service/source/tool"
)

var logger *logrus.Logger

func init() {
	logger = source.GetLogger()
}

func main() {
	logger.Infoln("server is running...")
	route := gin.New()
	route.Use(gin.Recovery(), middleware.LoggerMiddleware())
	router.RegisterRouter(route)
	err := route.Run(":10010")
	if err != nil {
		logger.Fatal(err)
	}
	env := tool.GetTool().GetEnv()
	config := source.Config
	if env == "product" {
		config.SetConfigName("config")
		logger.SetFormatter(&logrus.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	} else {
		config.SetConfigName("configTest")
		logger.SetFormatter(&logrus.TextFormatter{})
	}
	logger.Infoln("server exited...")
	mongoClient := source.GetMongoClient()
	_ = mongoClient.Disconnect(source.NewCtx())
	return
}
