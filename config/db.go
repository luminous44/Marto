package config

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
		db *gorm.DB
	)

	 func Connection() {
    d, err := gorm.Open("mysql", "root:29344@tcp(127.0.0.1:3306)/simple_db?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    db = d
    fmt.Println("Database connected successfully")
}
	
	func GetDB() *gorm.DB {
		return db
	}