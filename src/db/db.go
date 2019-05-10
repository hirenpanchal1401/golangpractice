package db

import (
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

//DbConn ...
func DbConn() (db *gorm.DB) {
	err1 := godotenv.Load()
	if err1 != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := gorm.Open(os.Getenv("DB_DIALECT"), os.Getenv("DB_USER")+":@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	} else {
		log.Println("database connection successful")
		return db
	}
	// defer db.Close()
}
