package db

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// initializes the database connection
func InitDB() {
	var err error

	db, err = gorm.Open("mysql", "root:1174Kavindu@@/todos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
}

// return the database instance
func GetDB() *gorm.DB {
	return db
}

//the log.Fatal function is used to log a message using the standard logger from the log package
