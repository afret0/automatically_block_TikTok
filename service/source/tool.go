package utils

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func GetGinEngine() *gin.Engine {
	return engine
}

func GetEnv() string {
	env := os.Getenv("environment")
	return env
}

func NowString() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
