package Config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

var (
	host     = GetEnv("PGHOST", "localhost")
	user     = GetEnv("PGUSER", "postgres")
	password = GetEnv("PGPASSWORD", "postgres")
	port     = GetEnv("PGPORT", "5432")
	dbname   = GetEnv("PGDATABASE", "db_go_nexmedis")
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
