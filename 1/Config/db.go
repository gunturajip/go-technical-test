package Config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	port     = os.Getenv("PGPORT")
	dbname   = os.Getenv("PGDATABASE")
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if sqlDB != nil {
		err := sqlDB.Close()
		if err != nil {
			log.Fatal("error closing to database :", err)
		}
	}
}

func GetDB() *gorm.DB {
	return db
}
