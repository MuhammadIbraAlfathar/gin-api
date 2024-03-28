package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func LoadDB() {
	connectionStr := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_URL, ENV.DB_DATABASE)
	db, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connection database")
	}

	DB = db

}
