package source

import (
	"context"
	"github.com/sirupsen/logrus"
	"service/source/tool"
	"time"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	env := tool.GetTool().GetEnv()
	if env == "pro" {
		logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{ForceColors: true, FullTimestamp: true})
	}
}

func GetLogger() *logrus.Logger {
	return logger
}

func NewCtx() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	return ctx
}
