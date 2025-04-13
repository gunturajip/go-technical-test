package main

import (
	"1/Config"
	"1/Model"
)

func main() {
	Config.StartDB()
	db := Config.GetDB()
	db.Debug().AutoMigrate(
		Model.User{},
		Model.Product{},
		Model.Order{},
		Model.Item{},
	)
	defer Config.CloseDB()
}
