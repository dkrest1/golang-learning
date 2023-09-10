package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func init() {
    if err := godotenv.Load(); err != nil {
        panic("Error loading .env file")
    }
}

var (
	db *gorm.DB
)

func  Connect() {
	d, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}
	
	db = d
}

func GetDB () * gorm.DB {
	return db
}