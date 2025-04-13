package main

import (
	"1/Config"
	"1/Router"
	"os"
)

func main() {
	Config.StartDB()
	defer Config.CloseDB()
	r := Router.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
