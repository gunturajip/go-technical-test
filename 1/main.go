package main

import (
	"1/Config"
	"1/Router"
)

func main() {
	Config.StartDB()
	defer Config.CloseDB()
	r := Router.StartApp()
	var PORT = Config.GetEnv("PORT", "8080")
	r.Run(":" + PORT)
}
