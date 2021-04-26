package main

import (
	"backend/router"
	"backend/source"
)

func main() {
	source.Logger.Info("server is running...")
	E := source.GetGinEngine()
	router.RegisterRouter()
	err := E.Run("127.0.0.1:10010")
	if err != nil {
		source.Logger.Fatal(err)
	}
	source.Logger.Info("server exited...")
	return
}
