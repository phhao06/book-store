package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	e := godotenv.Load(filepath.Join("/home/haopham/workspace/go-projects/book-store/", "local.env"))
	if e != nil {
		log.Fatal("Error when loading env file \n", e)
	}
	username := os.Getenv("username")
	password := os.Getenv("password")
	database := os.Getenv("database")
	DBServer := os.Getenv("DBServer")
	dsn := username + ":" + password + "@tcp(" + DBServer + ":3306)/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//d, err := gorm.Open("mysql", "username:password/database?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Println("Error when connecting to MYSQL server")
		panic(err)
	}
	fmt.Println("Connected")
	db = d
}

func GetDB() *gorm.DB {
	return db
}
