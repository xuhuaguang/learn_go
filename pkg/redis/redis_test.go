package redis

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"learn_go/models"
	"learn_go/utils"
	"math/rand"
	"testing"
	"time"
)

var poolOption = &PoolOption{
	Address:     "127.0.0.1:6379",
	Password:    "",
	DB:          0,
	MaxIdle:     10,
	MaxActive:   10,
	IdleTimeout: 10,
}

func TestNotFound(t *testing.T) {
	connPool, _ := NewPool(poolOption)
	s, e := connPool.GetString("a-123")

	if e != nil {
		println(s)
	} else {
		println(e.Error())
	}

	if e == redis.ErrNil {
		println("not found")
	}

	if e == NotFoundErr {
		println("not found2")
	}
}

func TestConnPool_Hash(t *testing.T) {
	connPool, _ := NewPool(poolOption)

	_, _ = connPool.HSet("user:1", "name", "donghai")
	hGetAll, _ := connPool.HGetAll("user:1")

	json1, _ := json.Marshal(hGetAll)
	println(string(json1))
	strings, _ := redis.StringMap(hGetAll, nil)
	json2, _ := json.Marshal(strings)
	println(string(json2))

	name, _ := connPool.HGetString("user:1", "name")
	println(name)
}

func TestConnPool_HGetAllStringMap(t *testing.T) {
	connPool, _ := NewPool(poolOption)
	stringMap, e := connPool.HGetAllStringMap("12345")
	if e != nil {
		println(e.Error())
	} else {
		println(len(stringMap))
	}

}

func TestConnPool_IncrBy(t *testing.T) {
	connPool, _ := NewPool(poolOption)
	i, e := connPool.HIncrBy("count1", "imp", 1)
	if e == nil {
		println(i)
	} else {
		println(e.Error())
	}
}

func TestConnPool_HMGetStringMap(t *testing.T) {
	connPool, _ := NewPool(poolOption)
	stringMap, e := connPool.HMGetStringMap("12345", "12", "2039")
	if e != nil && e != redis.ErrNil {
		println(e.Error())
	} else {
		println(len(stringMap))
	}

}

func TestFreqPool_IncrBy(t *testing.T) {
	date := time.Now().Format(models.DateFormat_yyyyMMdd)
	usr := "i123456789"
	sid := "111666"
	key := fmt.Sprintf("freq:%s:%s", date, usr)
	newSid := fmt.Sprintf("%s_time", sid)
	newSidCount := fmt.Sprintf("%s_count", sid)
	connPool, _ := NewPool(poolOption)
	_, _ = connPool.HSet(key, newSid, time.Now().Unix())
	_, _ = connPool.HIncrBy(key, newSidCount, 1)

	sidArray := []string{"111222_time", "111333_time", "111444_time"}
	var timeArray []interface{}
	for i := 0; i < 3; i++ {
		minute := utils.RandInt(30, 60) + 30
		timestamp := utils.AddMinuteTimestamp(minute)
		timeArray = append(timeArray, timestamp)
	}
	_, _ = connPool.HMSet(key, sidArray, timeArray)

	_, _ = connPool.HSet(key, "tanx_screen", utils.GetAddAfterTime(models.MinMinute, models.MaxMinute, models.FixedMinute))
}

func TestTimers(t *testing.T) {
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Second)
		i := utils.RandInt(models.MinMinute, models.MaxMinute) + models.FixedMinute
		println(i)
	}
	timestamp := utils.GetAddAfterTime(models.MinMinute, models.MaxMinute, models.FixedMinute)
	println(timestamp)

}
func TestFreqPool_GET(t *testing.T) {
	date := time.Now().Format(models.DateFormat_yyyyMMdd)
	usr := "i123456789"
	key := fmt.Sprintf("freq:%s:%s", date, usr)
	connPool, _ := NewPool(poolOption)
	hGetAll, _ := connPool.HGetAll(key)
	strings, _ := redis.StringMap(hGetAll, nil)
	json2, _ := json.Marshal(strings)
	println(string(json2))

	name, _ := connPool.HGetString(key, "111666_time")
	println(name)

	stringMap, e := connPool.HMGetStringMap(key, "111666_time", "111666_count")
	if e != nil && e != redis.ErrNil {
		println(e.Error())
	} else {
		println(len(stringMap))
	}

	stringMapAll, e := connPool.HGetAllStringMap(key)
	if e != nil {
		println(e.Error())
	} else {
		println(len(stringMapAll))
	}
}

func TestActivity_IncrBy(t *testing.T) {
	date := time.Now().Format(models.DateFormat_yyyyMMdd)
	usr := "i123456789"
	sid := "1"
	key := fmt.Sprintf("ac:%s:%s", date, usr)
	connPool, _ := NewPool(poolOption)
	//_, _ = connPool.HSet(key, sid, time.Now().Unix())
	_, _ = connPool.HIncrBy(key, sid, 1)
	_, _ = connPool.ExpireKey(key, 24*60*60)
}

//测试特殊的tanx
func TestTanx(t *testing.T) {
	for i := 0; i <= 50; i++ {
		time.Sleep(3000 * time.Millisecond)
		GenerateRangeNum(0, 100)
	}
}

// GenerateRangeNum 生成一个区间范围的随机数
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

func TestUsrMd5EndValue(t *testing.T) {
	connPool, _ := NewPool(poolOption)
	setString, e := connPool.SetString("usr_level_value", 2)
	fmt.Println(setString, e)
	fmt.Println("------------")
	value, err := connPool.GetString("usr_level_value")
	if err == nil {
		fmt.Println(value)
	}
}
