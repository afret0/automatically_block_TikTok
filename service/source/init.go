package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var engine *gin.Engine
var Logger *logrus.Logger
var middlewareL *logrus.Logger
var Config *viper.Viper
var ctx context.Context
var cancel context.CancelFunc

func init() {
	engine = gin.New()
	Logger = logrus.New()
	middlewareL = logrus.New()
	Config = viper.New()

	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetReportCaller(true)

	engine.Use(gin.Recovery(), LoggerMiddleware())

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	env := GetEnv()
	if env == "product" {
		Config.SetConfigName("config")
		engine.Use(gin.Recovery())
		Logger.SetFormatter(&logrus.JSONFormatter{})
		middlewareL.SetFormatter(&logrus.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	} else {
		Config.SetConfigName("configTest")
		Logger.SetFormatter(&logrus.TextFormatter{})
		middlewareL.SetFormatter(&logrus.TextFormatter{})
	}

	Config.AddConfigPath("./service/config")
	err := Config.ReadInConfig()
	if err != nil {
		Logger.Fatal(err)
	}

	DB = getDatabase()
}
