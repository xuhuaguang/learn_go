package learn_go

import (
	"encoding/json"
	"errors"
	"fmt"
	gocache "github.com/patrickmn/go-cache"
	"learn_go/entity"
	"runtime"
	"sync"
	"testing"
	"time"
)

const (
	OneDayTime       = 24 * time.Hour
	defaultCacheTime = 65 * time.Second
	DefaultMaxDspQps = float32(10000)

	NoExpiration      time.Duration = -1
	DefaultExpiration time.Duration = 0
)

var (
	ErrNotFound = errors.New("item not found")
)

var LogStatusCache = gocache.New(gocache.NoExpiration, gocache.NoExpiration)

func init() {
	// 默认这些接口的写日志，其他接口可以通 admin 管理接口控制写日志
	LogStatusCache.Set("uri:zy", true, gocache.NoExpiration)
	LogStatusCache.Set("uri:client", true, gocache.NoExpiration)
}

func SetLogStatus(uri string) {
	LogStatusCache.Set(uri, true, gocache.NoExpiration)
}

func GetLogStatus() map[string]bool {
	statusMap := make(map[string]bool)
	println("-----------------", LogStatusCache.Items())
	for k, v := range LogStatusCache.Items() {
		statusMap[k] = v.Object.(bool)
	}
	return statusMap
}

//缓存结构体
type CacheData struct {
	AcceptInfo *gocache.Cache
}

//默认赋值结构体
func NewCacheData() *CacheData {
	return &CacheData{AcceptInfo: gocache.New(OneDayTime, OneDayTime)}
}

//对内存设置数值
func (c *CacheData) SetAcceptInfo(accept *entity.Accept) {
	c.AcceptInfo.Set(fmt.Sprintf("%d", accept.Id), accept, 100*OneDayTime)
}

//获取内存中的数值
func (c *CacheData) GetAcceptInfo(acceptId int) (*entity.Accept, error) {
	if data, exist := c.AcceptInfo.Get(fmt.Sprintf("%d", acceptId)); exist {
		accept := data.(*entity.Accept)
		return accept, nil
	} else {
		return nil, ErrNotFound
	}
}

func TestCacheInfo(t *testing.T) {
	SetLogStatus("ad")
	cacheData := NewCacheData()
	accept := &entity.Accept{
		Id:          1,
		Name:        "创越",
		ServiceType: "1",
		Type:        "1",
		AppFields:   "1",
		Url:         "www.baidu.com",
		PriceKey:    "123456",
		Token:       "6666!@#",
		QpsLimit:    0,
		Code:        "Screen",
		LogoUrl:     "www.baidu.com",
	}
	cacheData.SetAcceptInfo(accept)

	time.Sleep(1000)
	acceptInfo, e := cacheData.GetAcceptInfo(accept.Id)
	if e != nil {
		fmt.Println("error=", e)
	}
	bytes, _ := json.Marshal(acceptInfo)
	fmt.Println(string(bytes))

	fmt.Println("---------------------------------")
	logStatus := GetLogStatus()
	fmt.Println(len(logStatus))
}

func TestGoCache(t *testing.T) {
	//创建一个默认过期时间为5分钟的缓存，并且
	//每10分钟清除过期项目
	c := gocache.New(5*time.Minute, 10*time.Minute)

	// 将键“foo”的值设置为“bar”，使用默认的过期时间
	c.Set("foo", "bar", gocache.DefaultExpiration)

	//将键“baz”的值设置为42，没有过期时间
	//在重新设置或使用删除之前，不会删除该项
	//c.Delete（“baz”）
	c.Set("baz", 42, gocache.NoExpiration)

	// 从缓存中获取与键“foo”相关联的字符串
	foo, found := c.Get("foo")
	if found {
		fmt.Println(foo)
	}

	//如果在同一个函数中多次使用该值，这会变得很乏味。
	//您可以执行以下任一操作：
	if x, ok := c.Get("foo"); ok {
		_ = x.(string) //字符串类型
	}

	// or
	var foo2 string
	if x, ok := c.Get("foo"); ok {
		foo2 = x.(string) //字符串类型
	}
	fmt.Println(foo2)

	//结构体
	c.Set("foo", &MyStruct{}, gocache.DefaultExpiration)
	if x, found := c.Get("foo"); found {
		foo3 := x.(*MyStruct)
		fmt.Println(foo3)
	}

}

type MyStruct struct {
}

// ----------------------------华丽的分割线-----------------------------

//Item foo already exists
//bar true
//map[foo:{bar 0}]
//仅当给定的项不存在时或者如果现有项已过期，才将项添加到缓存中，
//否则返回错误。
func TestAdd(t *testing.T) {
	cache := gocache.New(gocache.DefaultExpiration, 0)
	err := cache.Add("foo", "bar", gocache.DefaultExpiration)
	if err != nil {
		t.Error("Couldn't add foo even though it shouldn't exist")
	}
	err = cache.Add("foo", "baz", gocache.DefaultExpiration)
	if err == nil {
		t.Error("Successfully added another foo when it should have returned an error")
	} else {
		fmt.Println(err)
	}
	fmt.Println(cache.Get("foo"))
	fmt.Println(cache.Items())
}

//仅当缓存键已存在时才为其设置新值，并且
//项目尚未过期。否则返回错误。
//Item foo doesn't exist
//bird true
func TestReplace(t *testing.T) {
	tc := gocache.New(DefaultExpiration, 0)
	err := tc.Replace("foo", "bar", DefaultExpiration)
	if err == nil {
		t.Error("Replaced foo when it shouldn't exist")
	} else {
		fmt.Println(err)
	}
	tc.Set("foo", "bar", DefaultExpiration)
	err = tc.Replace("foo", "bird", DefaultExpiration)
	if err != nil {
		t.Error("Couldn't replace existing key foo")
	}
	fmt.Println(tc.Get("foo"))
}

//从缓存中删除项目。如果密钥不在缓存中，则不执行任何操作。
func TestDelete(t *testing.T) {
	tc := gocache.New(DefaultExpiration, 0)
	tc.Set("foo", "bar", DefaultExpiration)
	tc.Delete("foo")
	x, found := tc.Get("foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
	fmt.Println(x, found)
}

//返回缓存中的项数。这可能包括过期，但尚未清理。
func TestItemCount(t *testing.T) {
	tc := gocache.New(DefaultExpiration, 0)
	tc.Set("foo", "1", DefaultExpiration)
	tc.Set("bar", "2", DefaultExpiration)
	tc.Set("baz", "3", DefaultExpiration)
	if n := tc.ItemCount(); n != 3 {
		t.Errorf("Item count is not 3: %d", n)
	} else {
		fmt.Println(n)
	}
}

//Delete all items from the cache.
func TestFlush(t *testing.T) {
	tc := gocache.New(DefaultExpiration, 0)
	tc.Set("foo", "bar", DefaultExpiration)
	tc.Set("baz", "yes", DefaultExpiration)
	tc.Flush()
	x, found := tc.Get("foo")
	fmt.Println(x, found)
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
	x, found = tc.Get("baz")
	fmt.Println(x, found)
	if found {
		t.Error("baz was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
}

func TestIncrementOverflowInt(t *testing.T) {
	tc := gocache.New(DefaultExpiration, 0)
	tc.Set("int8", int8(127), DefaultExpiration)
	err := tc.Increment("int8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	} else {
		fmt.Println(err)
	}
	x, _ := tc.Get("int8")
	int8 := x.(int8)
	if int8 != -128 {
		t.Error("int8 did not overflow as expected; value:", int8)
	}
	fmt.Println(x, int8)
}

func BenchmarkRWMutexMapGetConcurrent(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}      //创建一把读写锁
	wg := new(sync.WaitGroup) //创建一个同步等待组的对象
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers) //设置同步等待组的数量
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				mu.RLock()
				_, _ = m["foo"]
				mu.RUnlock()
			}
			wg.Done() //执行完成 同步等待数量减1
		}()
	}
	wg.Wait() //主goroutine进入阻塞状态
}
