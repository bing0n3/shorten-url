package model

import (
	"log"

	"github.com/bing0n3/shorten-url/config"
	"github.com/jinzhu/gorm"

	// import mysql drive
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var counter int64

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Printf("Connet to db, url is \"%s\"...\n", connectingStr)
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}

// create table if table not exist
func CreateTable() {
	if !db.HasTable(&URL{}) {
		log.Println("URL table doesn't exist. Create...")
		db.CreateTable(&URL{})
	}
}
