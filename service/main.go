package main

import (
	"backend/router"
	"backend/source"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = source.GetLogger()
}

func main() {
	logger.Infoln("server is running...")
	E := source.GetGinEngine()
	router.RegisterRouter()
	err := E.Run("127.0.0.1:10010")
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infoln("server exited...")
	defer source.GetCancel()()
	return
}
