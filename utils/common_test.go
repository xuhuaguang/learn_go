package utils

import (
	"fmt"
	"github.com/segmentio/ksuid"
	"learn_go/entity"
	"learn_go/models"
	"math/rand"
	"net/url"
	"sort"
	"strings"
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
	a := []int{4, 2, 3, 1, 10, 5, 6, 7, 9, 8, 1, 2}
	loop := RemoveRepByLoop(a)
	fmt.Println(loop)
	sort.Ints(loop)
	fmt.Println(loop)

}

func TestIsContainsAnyValue(t *testing.T) {
	clickTrackers := []string{"http://et.w.inmobi.cn/c.asm/", "http://i.l.inmobicdn.cn/adtools", "http://c.gdt.qq.com"}
	contains := StrListValveContains(clickTrackers, "c.gdt.qq.com")
	fmt.Println(contains)
}

func TestRandNumber(t *testing.T) {
	//[0,8]返回true ,反之返回false
	num_true := 0
	num_false := 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		rule := RandShowRule(0, 100, RandInt(5, 10))
		if rule {
			num_true++
		} else {
			num_false++
		}
	}
	fmt.Println("true num ", num_true)
	fmt.Println("num_false num ", num_false)

	for i := 0; i < 10; i++ {
		fmt.Println(RandInt(5, 10))
	}

}

func CreateCaptcha() string {
	return fmt.Sprintf("%08v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
}

type Content struct {
	RawRequest interface{}
}

func TestInterfaceGetValue(t *testing.T) {
	info := &entity.RequestInfo{
		Date:      0,
		Hour:      0,
		EventTime: 0,
		ID:        ksuid.New().String(),
		PageNum:   12,
	}
	request := &Content{}
	request.RawRequest = info
	fmt.Println("--------------")
	requestInfo := request.RawRequest.(*entity.RequestInfo)
	fmt.Println(requestInfo.PageNum)
}

func TestMd5(t *testing.T) {
	_ = "zy@987qwer"
	signMap := map[string]interface{}{
		"usr":        "i1234556",
		"bookId":     "3",
		"expiredDay": 2,
		"source":     2,
		"p1":         "v_p1",
		"p2":         "v_p2",
		"p3":         "v_p3",
		"p4":         "v_p4",
		"p15":        "v_p15",
		"ip":         "v_ip",
		"_php_":      "SSP",
	}
	var keys []string
	for k := range signMap {
		keys = append(keys, k)
	}
	//按字典升序排列
	sort.Strings(keys)
	str := ""
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", signMap[k])
		str += fmt.Sprintf("%s%v", k, signMap[k])
	}
	fmt.Println(str)
	fmt.Println(Md5(str))
}

func TestStringsSplit(t *testing.T) {
	str := "5710,5730,5740,5741,5742,5743,"
	if strings.HasSuffix(str, ",") {
		str = str[:len(str)-1]
	}
	split := strings.Split(str, ",")
	fmt.Println(len(split))
}

func TestArrayData(t *testing.T) {
	keys := []string{"1", "2", "3", "4"}
	fmt.Println(checkArray(keys))

	ids := ""
	for _, v := range keys {
		ids += fmt.Sprintf("%s_", v)
	}
	fmt.Println(ids)
}

func checkArray(array []string) bool {
	for _, v := range array {
		if v == "3" {
			return false
		}
	}
	return true
}

func TestArrayAppend(t *testing.T) {
	keys := []string{"1", "2", "3", "4"}
	keys2 := []string{"1", "2", "3", "4", "5", "7", "3", "4", "12"}
	var allArray []string
	allArray = append(allArray, keys...)
	allArray = append(allArray, keys2...)

	fmt.Println(len(allArray))
	fmt.Println("-------------")
	for _, v := range allArray {
		fmt.Println(v)
	}

	allArray = DuplicateStringArray(allArray)
	fmt.Println("-------------")
	fmt.Println(len(allArray))
	if StrListContains(allArray, "3") {
		fmt.Println("包含--")
	}

	var allArray2 []string
	if StrListContains(allArray2, "3") {
		fmt.Println("包含--")
	}

	num, num2 := 1, 2
	if num == 1 || num2 == 4 {
		fmt.Println("num")
	}
}

func TestParseTime(t *testing.T) {
	times := 60034
	println(TransferTime(int32(times)))

	id := GetNewPlanMappingId("zybc0e74")
	fmt.Println(id)
	println(GetPlanMappingId(id))
}

func TestRandIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		index := rand.Intn(5)
		fmt.Println(index)
	}

	fmt.Println(3 < 3)
}

func TestContent(t *testing.T) {
	timeout := time.Duration(300) * time.Millisecond
	println(timeout)
}

func TestDateNum(t *testing.T) {
	for i := 0; i < 5; i++ {
		date := time.Now().AddDate(0, 0, -i)
		dt := date.Format(models.DateFormat_yyyyMMdd)
		println(dt)
	}
}

//获取url的域名
func TestUrlDomain(t *testing.T) {
	str := "https://fc.ele.me/a/ODE0NDg2MWI0YjFjMTFlYjlhM2QwMjQyMGI1OWUxMjQ="
	domain := strings.Split(strings.Split(str, "//")[1], "/")[0]
	println(domain)
	s, _ := url.Parse("https://fc.ele.me/a/ODE0NDg2MWI0YjFjMTFlYjlhM2QwMjQyMGI1OWUxMjQ=")
	println(s.Host)
}

func TestCidMd5(t *testing.T) {
	resources := []string{"https://img1.rgyun.net.cn/adxupmat/2021/08/03/16cfe944-ee3e-416a-b93d-4d0c477654c5.jpg"}
	ids := fmt.Sprintf("%d:%d:%v", 10012, 10052, resources)
	println(Md5(ids))
}

func TestJitter(t *testing.T) {
	for i := 0; i < 1; i++ {
		wait := Jitter(1, 0.25)
		println(wait)
	}
	wait := Jitter(10, 0.25)
	println(wait)
}

func TestDate(t *testing.T) {
	for i := 0; i < 5; i++ {
		date := time.Now().AddDate(0, 0, -i)
		println(date.Format(models.DateFormat_yyyyMMdd))
	}
}

/**
  1、f1=true  f2=false  f3=true    -->true
  2、f1=true  f2=false  f3=false   -->false x
  3、f1=false f2=true   f3=false   -->false
  4、f1=false f2=true   f3=true    -->true
  5、f1=false  f2=false  f3=true   -->false
  6、f1=false  f2=false  f3=false  -->false
*/
func TestIf(t *testing.T) {
	f1 := true
	f2 := false
	f3 := false

	if f1 || f2 && f3 {
		println("------true")
	} else {
		println("=====false")
	}
}
