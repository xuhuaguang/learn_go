package learn_go

import (
	"fmt"
	"go.uber.org/atomic"
	"testing"
)

// 标准库的sync/atomic功能强大， 但很容易忘记哪些变量必须原子地访问。
// go.uber.org/atomic保留标准库的所有功能，但封装基本类型以提供更安全、更方便的API。

func TestExample(t *testing.T) {
	var atom atomic.Int32
	atom.Store(42)
	fmt.Println(atom.Inc())      //原子加1 ==>43
	fmt.Println(atom.CAS(43, 0)) //比较并交换 ==>true
	fmt.Println(atom.Load())     //获取原子值 ==>0
}

//其中：新增=ADD、获取=GET、相减=SUB
func TestInt32AllMethod(t *testing.T) {
	at := atomic.NewInt32(40) //创建默认值  --40
	fmt.Println(at.Add(2))    //ADD 原子方式添加指定的数值到int32,并返回新值  --42
	fmt.Println(at.Load())    //GET 原子方式加载获取包装的值 --42
	fmt.Println(at.Sub(10))   //SUB 从包装的int32中减法并返回新值。--32
	fmt.Println(at.Inc())     //ADD 原子加1，内部调用的是at.Add(1)的方法 --33
	fmt.Println(at.Dec())     //SUB Dec以原子方式递减并返回新值。内部调用的是at.Sub(1)的方法--32

	fmt.Println(at.CAS(32, 99)) //CAS是一种原子比较和交换。old值是当前原子的值，new是新的赋值。如果填写的old值等于当前原子值，那就返回true
	fmt.Println(at.Load())      ///result: 99
	at.Store(88)                //ADD Store以原子方式存储传递的值。
	fmt.Println(at.Load())      //result: 88

	fmt.Println(at.Swap(8))  //Swap原子交换包装好的int32并返回旧值。result: 88
	fmt.Println(at.Load())   //result: 8
	fmt.Println(at.String()) //result:8 字符串

	bytes, _ := at.MarshalJSON()
	fmt.Println(bytes)
	fmt.Println(at.UnmarshalJSON(bytes))
}
