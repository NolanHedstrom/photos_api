package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

//DB is how you refernce the database
var DB *gorm.DB

//OpenDB connects to the DB
func OpenDB(databaseConn string, mode bool) error {

	DB, err := gorm.Open("mysql", databaseConn)
	if err != nil {
		log.Fatal("Error in connecting to the database: ", err.Error())
		return err
	}

	DB.LogMode(mode)

	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	return nil
}

//CloseDB closes the DB connection
func CloseDB(db *gorm.DB) {
	db.Close()
}
