package utils

import (
	"fmt"
	"testing"
)

func TestGetRedisOption(t *testing.T) {
	InitConfig()
	ops := GetRedisOption()
	fmt.Println(ops.Addr, ops.DB, ops.PoolSize)
}