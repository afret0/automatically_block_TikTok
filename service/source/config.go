package source

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"service/source/tool"
)

var Config *viper.Viper

func init() {
	Config = viper.New()
	Config.AddConfigPath("./config")
	env := tool.GetTool().GetEnv()
	if env == "product" {
		Config.SetConfigName("config")
		gin.SetMode(gin.ReleaseMode)
	} else {
		Config.SetConfigName("configTest")
	}
	err := Config.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
}

func GetConfig() *viper.Viper {
	return Config
}
