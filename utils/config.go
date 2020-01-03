package utils

import (
	"fmt"
	"github.com/go-redis/redis/v7"

	"github.com/spf13/viper"
)

var (
	projectName string = "shorten-url"
)

func InitConfig() {
	getConfig(projectName)
}

func getConfig(projectName string) {
	viper.SetConfigName("shorten") // name of configs file (without extension)

	viper.AddConfigPath(".")                                                 // optionally look for configs in the working directory
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s", projectName))               // call multiple times to add many search paths
	viper.AddConfigPath(fmt.Sprintf("/data/docker/configs/%s", projectName)) // path to look for the configs file in
	viper.AddConfigPath("../../configs")
	viper.AddConfigPath("../configs")

	err := viper.ReadInConfig() // Find and read the configs file
	if err != nil {             // Handle errors reading the configs file
		panic(fmt.Errorf("Fatal error configs file: %s", err))
	}
}

func GetMysqlConnectingString() string {
	usr := viper.GetString("mysql.user")
	pwd := viper.GetString("mysql.password")
	host := viper.GetString("mysql.host")
	db := viper.GetString("mysql.db")
	charset := viper.GetString("mysql.charset")
	port := viper.GetInt("mysql.port")
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=%s&parseTime=true", usr, pwd, host, port, db, charset)
}

func GetRedisOption() redis.Options {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	db := viper.GetInt("redis.db")
	pool_size := viper.GetInt("redis.pool_size")

	return redis.Options{Addr: host + ":" + port, DB: db, PoolSize: pool_size}
}
