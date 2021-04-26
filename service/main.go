package main

import (
	"backend/router"
	"backend/utils"
)

func main() {
	utils.Logger.Info("server is running...")
	E := utils.GetGinEngine()
	router.RegisterRouter()
	err := E.Run("127.0.0.1:10110")
	if err != nil {
		utils.Logger.Fatal(err)
	}
	utils.Logger.Info("server exited...")
	return
}
