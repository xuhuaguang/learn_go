package learn_go

import (
	"fmt"
	"testing"
)

// 接口的意义是对其他类型的一个概括，接口内可以定义很多个方法，谁将这些方法实现， 就可以认为是实现了该接口。
// Go语言的多态，主要是通过接口来实现。

func Test_Interface(t *testing.T) {

	h := JiaHuoHao{name: "好家伙"}
	fmt.Println(h.name)

	l := LaoLiTou{name: "老李头"}
	fmt.Println(l.name)

	m := MyselfTest{name: "自己测试名称"}
	testInterface(m)
	//testInterface 需要参数类型为gongFu接口类型的参数
	//h实现了gongFu接口的方法
	//h就是这个接口的实现 就可以作为这个函数的参数
	testInterface(h)

	var kf gongFu //接口声明
	kf = h
	kf.Toad()

	l.Toad()
	l.PlayGame()
}

//测试方法
func testInterface(k gongFu) {
	k.Toad()
	k.SixSwords()
}

//定义接口
type gongFu interface {
	Toad()      //蛤蟆功
	SixSwords() //六脉神剑
}

//实现类
type JiaHuoHao struct {
	name string
}

//实现类
type LaoLiTou struct {
	name string
}

//自己测试实现类
type MyselfTest struct {
	name string
}

//实现方法JiaHuoHao --蛤蟆功
func (o JiaHuoHao) Toad() {
	fmt.Println(o.name, "实现了蛤蟆功..")
}

//实现方法JiaHuoHao --六脉神剑
func (o JiaHuoHao) SixSwords() {
	fmt.Println(o.name, "实现了六脉神剑..")
}

//实现方法LaoLiTou --蛤蟆功
func (f LaoLiTou) Toad() {
	fmt.Println(f.name, "也实现了蛤蟆功..")
}

//实现方法LaoLiTou --六脉神剑
func (f LaoLiTou) SixSwords() {
	fmt.Println(f.name, "也实现了六脉神剑.")
}

//实现自己的方法
func (f LaoLiTou) PlayGame() {
	fmt.Println(f.name, "玩游戏..")
}

//自己测试实现方法--新建的结构体必须实现全部接口中的方法
func (m MyselfTest) Toad() {
	fmt.Println(m.name, "实现了蛤蟆功..")
}

//自己测试实现方法--新建的结构体必须实现全部接口中的方法
func (m MyselfTest) SixSwords() {
	fmt.Println(m.name, "实现了六脉神剑..")
}

// 使用接口对方法进行约束，然后让方法实现接口，这样规范了方法。
// 通过使用同样的接口名称，但是在调用的时候使用不同的类，实现执行不同的方法。
// 这样就实现了Go语言中的多态。

// -------------------------------------------------以下是空接口---------------------------------------------------------

// 空接口就是不包含任何方法的接口，所有的类型都可以实现空接口，
// 因此空接口可以实现存储任意类型的数据， 谁实现它就被看作是谁的实现类。
func Test_Empty_Interface(t *testing.T) {
	EmptyTest(0)       //使用int类型做为参数传入到函数
	EmptyTest("ceshi") //使用string类型做为参数传入到函数

	EmptyTest2(0)       //使用int类型做为参数传入到函数
	EmptyTest2("ceshi") //使用string类型做为参数传入到函数
}

//空接口
type T interface{}

//定义一个函数 接收EmptyTest接口类型的数据
func EmptyTest(t T) {
	fmt.Println(t)
}

//函数简化版
func EmptyTest2(t interface{}) {
	fmt.Println(t)
}

// 可以将空接口类型写为interface{} 这种类型可以理解为任何类型。类似其他语言中的object。

// -------------------------------------------------以下是空接口实现------------------------------------------------------

//空接口既然可以传任意类型，利用这个特性可以把空接口interface{}当做容器使用
func Test_Empty_Use_Interface(t *testing.T) {
	//创建一个map类型 key为string val为空接口，这样值就可以存储任意类型了
	m := make(map[string]interface{})
	m["a"] = "zhangsan"
	m["b"] = 1.1
	m["c"] = true
	fmt.Println(m)

	fmt.Println("--------------------------------")
	// 创建字典实例
	dict := NewDict()
	// 添加数据
	dict.SetData("001", "第一条数据")
	dict.SetData("002", 3.1415)
	dict.SetData("003", false)
	// 获取值
	d := dict.GetData("001")
	fmt.Println(d)
	d2 := dict.GetData("002")
	fmt.Println(d2)
	d3 := dict.GetData("003")
	fmt.Println(d3)
}

// 字典结构
type Dictionary struct {
	data map[string]interface{} // 数据key为string值为interface{}类型
}

// 获取值
func (d *Dictionary) GetData(key string) interface{} {
	return d.data[key]
}

// 设置值
func (d *Dictionary) SetData(key string, value interface{}) {
	d.data[key] = value
}

// 创建一个字典
func NewDict() *Dictionary {
	return &Dictionary{
		data: make(map[string]interface{}), //map类型使用前需要初始化，所以需要使用make创建 防止空指针异常。
	}
}
