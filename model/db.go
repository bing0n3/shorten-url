package model

import (
	"log"
	"sync"

	"github.com/bing0n3/shorten-url/config"
	"github.com/jinzhu/gorm"

	// import mysql drive
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Counter struct {
	sync.Mutex
	count int64
}

var (
	db     *gorm.DB
	lastID *Counter
)

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

//setCounter func
func SetCounter(id *Counter) {
	lastID = id
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

// create table if table not exist func
func CreateTable() {
	if !db.HasTable(&Shorten{}) {
		log.Println("Shorten table doesn't exist. Create...")
		db.CreateTable(&Shorten{})
	}
}

// init
func InitLastID() *Counter {
	id, err := GetCounter()
	log.Printf("Got counter: %d\n", id)
	if err != nil {
		log.Println(err)
		panic("Cannot get counter")
	}
	counter := Counter{count: id}
	return &counter
}

// update lastID, and return a new id to save
// self-increasement is atom operate
func (counter *Counter) UpdateCounter() int64 {
	counter.Lock()
	counter.count++
	counter.Unlock()
	return counter.count
}
