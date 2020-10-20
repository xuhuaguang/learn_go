package learn_go

import (
	"encoding/json"
	"fmt"
	"github.com/chilts/sid"
	"github.com/kjk/betterguid"
	"github.com/oklog/ulid"
	"github.com/rs/xid"
	"github.com/satori/go.uuid"
	"github.com/segmentio/ksuid"
	"github.com/sony/sonyflake"
	"learn_go/utils"
	"log"
	"math/rand"
	"testing"
	"time"
)

//一、请求唯一id是可以转换成时间戳形式，可做曝光去重
// 步骤：
//1、请求到素材时，在redis中记录类型是hash。
//      key ==>  前缀_reqId  ( info.ksuid.String() )
//      value ==> 时间戳
//
//2、在曝光监测中，已经把当时的时间戳拼接在url中了（假设：prev_time=URL时间串）
//      然后通过 ksuid.Parse(ReqId)获取此请求id时的时间戳，再和prev_time比对，防止非法的请求。
//
//3、redis的作用是校验 曝光事件，检查是否存在竞价。因为有竞价的曝光才有意义。
//      首次曝光之后会记录此次事件，防止重复曝光。

func TestKsuid(t *testing.T) {
	ksuids := utils.GenKsuid() //新建一个ksuids对象
	info := &TestKusidInfo{
		ksuid:    &ksuids,
		PrevTime: int64(ksuids.Timestamp()), //获取当前的时间戳
	}
	reqId := info.ksuid.String() //获取当前的uuid
	info.ReqId = reqId

	bytes, _ := json.Marshal(info)
	fmt.Println(string(bytes))

	parse, _ := ksuid.Parse(reqId) //解析uuid，获取当时ksuids对象
	info.ksuid = &parse
	fmt.Println(parse.Timestamp()) //获取当时的时间戳。和刚开始的PrevTime一样
}

type TestKusidInfo struct {
	ksuid    *ksuid.KSUID
	PrevTime int64
	ReqId    string
}

// *******************************************************************************************************************//

func genXid() {
	id := xid.New()
	fmt.Printf("github.com/rs/xid:           %s\n", id.String())
}

func genKsuid() {
	id := ksuid.New()
	fmt.Printf("github.com/segmentio/ksuid:  %s\n", id.String())
}

func genBetterGUID() {
	id := betterguid.New()
	fmt.Printf("github.com/kjk/betterguid:   %s\n", id)
}

func genUlid() {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	fmt.Printf("github.com/oklog/ulid:       %s\n", id.String())
}

func genSonyflake() {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	// Note: this is base16, could shorten by encoding as base62 string
	fmt.Printf("github.com/sony/sonyflake:   %x\n", id)
}

func genSid() {
	id := sid.Id()
	fmt.Printf("github.com/chilts/sid:       %s\n", id)
}

func genUUIDv4() {
	id := uuid.NewV4()
	fmt.Printf("github.com/satori/go.uuid:   %s\n", id)
}

//一个好的随机唯一ID遵循以下原则：
//	1、唯一性：基本原则，必须满足。
//	2、可排序的：可以使用随机唯一ID字符串进行排序。
//	3、具有时间属性：同一时间内，生成的ID彼此相近。
//	4、随机唯一ID字符串无需转义就要以做为URL的一部份。
//	5、越短越好
func TestUuidAllMethod(t *testing.T) {
	genXid()
	genKsuid()
	genBetterGUID()
	genUlid()
	genSonyflake()
	genSid()
	genUUIDv4()
}
