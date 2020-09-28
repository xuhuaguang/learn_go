package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestDownPtrInfoToLocal(t *testing.T) {
	fields := ptrRedisKeys(10)
	fmt.Println(fields)
}

func ptrRedisKeys(c int) []string {
	now := time.Now().Local()
	var keyFields []string
	for i := 0; i < c; i++ {
		t := now.Add(time.Duration(-i) * time.Minute)
		hkey1 := fmt.Sprintf("%d%d", t.Hour(), t.Minute())
		keyFields = append(keyFields, hkey1)
	}
	return keyFields
}

type SaleNumType int //数量限制类型 0=展示 1=请求,默认展示

const (
	SaleNumTypeImp SaleNumType = 0 //展示
	SaleNumTypeReq SaleNumType = 1 //请求
)

var SaleNumTypeMap = map[SaleNumType]string{
	SaleNumTypeImp: "imp",
	SaleNumTypeReq: "req",
}

func TestType(t *testing.T) {
	var num SaleNumType = 1
	s := SaleNumTypeMap[num]
	fmt.Println(s)
}
