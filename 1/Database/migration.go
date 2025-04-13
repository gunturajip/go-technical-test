package Database

import (
	"1/Config"
	"1/Model"
)

func main() {
	Config.StartDB()
	db := Config.GetDB()
	db.Debug().AutoMigrate(
		Model.User{},
		Model.Order{},
		Model.Item{},
		Model.Product{},
	)
	defer Config.CloseDB()
}
