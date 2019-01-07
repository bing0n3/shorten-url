package main

import (
	"github.com/bing0n3/shorten-url/controller"
	"github.com/bing0n3/shorten-url/model"
)

func main() {

	// connect to DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// create table if table not exist.
	model.CreateTable()
	counter := model.InitLastID()
	model.SetCounter(counter)

	controller.StartRouter()
}
