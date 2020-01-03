package mysql

import (
	"github.com/bing0n3/shorten-url/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"sync"
)


type MySQLClient struct {
	*sqlx.DB
}

var instance *MySQLClient
var once sync.Once


// GetMySQLClient get a RedisClient instance.
func GetMySQLClient() (*MySQLClient) {
	once.Do(func() {
		db, err := sqlx.Connect("mysql", utils.GetMysqlConnectingString())
		utils.Info.Printf("Connect with mysql by %s", utils.GetMysqlConnectingString())

		if err != nil {
			utils.Error.Panic("Cannot connect to mysql client!")
		}
		instance = &MySQLClient{db}
	})

	return instance
}

