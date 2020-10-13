package utils

import (
	"fmt"
	"sort"
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

func TestFloatCompare(t *testing.T) {
	v1 := 0.01
	v2 := 0.09

	i := int64(v1 * float64(100))
	i2 := int64(v2 * float64(100))
	fmt.Println(i)
	fmt.Println(i2)
}

func TestIntArraySort(t *testing.T) {
	arr := []int{3, 8, 2, 5, 1, 7, 9, 4}
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println("---------------------")
}

func TestDuplicateIntArray(t *testing.T) {
	a := []int{4, 2, 3, 1, 10, 5, 6, 7, 9, 8, 1, 2,}
	loop := RemoveRepByLoop(a)
	fmt.Println(loop)
	sort.Ints(loop)
	fmt.Println(loop)
}
