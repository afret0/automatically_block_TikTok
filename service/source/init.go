package source

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var engine *gin.Engine
var Logger *logrus.Logger
var middlewareLogger *logrus.Logger
var Config *viper.Viper
var ctx context.Context
var cancel context.CancelFunc

func init() {
	engine = gin.New()
	Logger = logrus.New()
	middlewareLogger = logrus.New()
	Config = viper.New()

	Logger.SetLevel(logrus.InfoLevel)
	Logger.SetReportCaller(true)

	engine.Use(gin.Recovery(), LoggerMiddleware())

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	env := GetEnv()
	engine.Use(gin.Recovery())
	if env == "product" {
		Config.SetConfigName("config")
		Logger.SetFormatter(&logrus.JSONFormatter{})
		middlewareLogger.SetFormatter(&logrus.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	} else {
		Config.SetConfigName("configTest")
		Logger.SetFormatter(&logrus.TextFormatter{})
		middlewareLogger.SetFormatter(&logrus.TextFormatter{})
	}

	Config.AddConfigPath("./config")
	err := Config.ReadInConfig()
	if err != nil {
		Logger.Fatal(err)
	}

	DB = getDatabase()
}
