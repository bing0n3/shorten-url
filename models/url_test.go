package models

import (
	"fmt"
	"github.com/bing0n3/shorten-url/utils"
	"testing"
)

func TestInsertAlias(t *testing.T) {
	utils.InitConfig()
	utils.InitLogDebug()
	err := InsertAlias(1,"baidu.com", "1h")
	if err != nil {
		t.Fatal(err)
	}
}


func TestGetURLByAlias(t *testing.T) {
	utils.InitConfig()
	utils.InitLogDebug()
	url, err := GetURLByAlias(1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(url)
}


func TestGetMaxAlias(t *testing.T) {
	utils.InitConfig()
	utils.InitLogDebug()
	alias, err := GetMaxAlias()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(alias)
}