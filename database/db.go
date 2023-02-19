package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	
	"gorm.io/gorm"
)

const DB_USERNAME = "yg44om958iie64y0"
const DB_PASSWORD = "r77jpsqasz2v392w"
const DB_NAME = "qo08agkjf6rg2k63"
const DB_HOST = "r77jpsqasz2v392w"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	ConnString := "yg44om958iie64y0:r77jpsqasz2v392w@tcp(qz8si2yulh3i7gl3.cbetxkdyhwsb.us-east-1.rds.amazonaws.com:3306)/qo08agkjf6rg2k63?parseTime=true&loc=Local"
	var err error
	dsn :=  ConnString
	fmt.Println("dsn :is workig ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}