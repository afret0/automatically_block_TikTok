package source

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var logger *logrus.Logger
var Config *viper.Viper

func init() {
	logger = logrus.New()
	Config = viper.New()
	Config.AddConfigPath("./config")
	err := Config.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}

	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	DB = getDatabase()
}

func GetLogger() *logrus.Logger {
	return logger
}

func NewCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	return ctx
}

func GetConfig() *viper.Viper {
	return Config
}
