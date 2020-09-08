package learn_go

/**
	map在Go语言并发编程中,如果仅用于读取数据时候是安全的，但是在读写操作的时候是不安全的，
    在Go语言1.9版本后提供了一种并发安全的，sync.Map是Go语言提供的内置map，不同于基本的map数据类型，
    所以不能像操作基本map那样的方式操作数据，他提供了特有的方法，不需要初始化操作实现增删改查的操作。
*/
import (
	"fmt"
	"sync"
	"testing"
)

//声明sync.Map
var syncmap sync.Map

func Test_Sync_Map(t *testing.T) {
	//Store方法将键值对保存到sync.Map
	syncmap.Store("zhangsan", 97)
	syncmap.Store("lisi", 100)
	syncmap.Store("wangmazi", 200)

	// Load方法获取sync.Map 键所对应的值
	fmt.Println(syncmap.Load("lisi"))

	// Delete方法键删除对应的键值对
	syncmap.Delete("lisi")

	// Range遍历所有sync.Map中的键值对
	syncmap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
