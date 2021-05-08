package source

import (
	"backend/source/tool"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var engine *gin.Engine
var logger *logrus.Logger
var middlewareLogger *logrus.Logger
var Config *viper.Viper
var ctx context.Context
var cancel context.CancelFunc

func init() {
	engine = gin.New()
	logger = logrus.New()
	middlewareLogger = logrus.New()
	Config = viper.New()

	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)

	engine.Use(gin.Recovery(), LoggerMiddleware())

	//ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)

	env := tool.GetEnv()
	engine.Use(gin.Recovery())
	if env == "product" {
		Config.SetConfigName("config")
		logger.SetFormatter(&logrus.JSONFormatter{})
		middlewareLogger.SetFormatter(&logrus.JSONFormatter{})
		gin.SetMode(gin.ReleaseMode)
	} else {
		Config.SetConfigName("configTest")
		logger.SetFormatter(&logrus.TextFormatter{})
		middlewareLogger.SetFormatter(&logrus.TextFormatter{})
	}

	Config.AddConfigPath("./config")
	err := Config.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}

	DB = getDatabase()
}

func GetLogger() *logrus.Logger {
	return logger
}

func GetCtx() context.Context {
	return ctx
}

func GetCancel() context.CancelFunc {
	return cancel
}
