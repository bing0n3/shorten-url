package mysql

import (
	"github.com/bing0n3/shorten-url/utils"
	"testing"
)

func TestGetMySQLClient(t *testing.T) {

	utils.InitLogDebug()
	utils.InitConfig()
	client := GetMySQLClient()
	err1 := client.Ping()
	if err1 != nil {
		t.Fatal("Failed to connect sql server")
	}
}