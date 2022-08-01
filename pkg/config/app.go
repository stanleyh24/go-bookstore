package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DSN = "root:admin123@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("DB connection established successfully")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
