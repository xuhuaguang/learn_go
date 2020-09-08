package learn_go

import (
	"fmt"
	"testing"
)

//1、定义结构体实现封装
type Haojiahuo struct {
	Name string
	Age  int
}

//使用NewPerson方法创建一个对象
func NewPerson(name string) *Haojiahuo {
	return &Haojiahuo{
		Name: name,
	}
}

// 使用SetAge方法设置结构体成员的Age
func (h *Haojiahuo) SetAge(age int) {
	h.Age = age
}

// 使用GetAge方法获取成员现在的Age
func (h *Haojiahuo) GetAge() int {
	return h.Age
}

func Test_Package(t *testing.T) {
	//创建一个对象
	h := NewPerson("好家伙")
	h.SetAge(18)                    //访问封装的方法设置年龄
	fmt.Println(h.Name, h.GetAge()) //使用对象封装的方法获取年龄
}

//2、继承的实现
//继承可以解决代码复用的问题，结构体内嵌套一个匿名结构体，也可以嵌套多层结构体

// 创建一个结构体起名 Ouyangcrazy 代表父类
type OuYangCrazy struct {
	Name    string
	Age     int
	Ability string
}

//创建一个结构体代表子类
type YangGuo struct {
	OuYangCrazy        //包含父类所有属性
	Address     string //单独子类有的字段
}

// 父类的方法
func (o *OuYangCrazy) ToadKongfu() {
	fmt.Println(o.Name, "的蛤蟆功！")
}

//子类的方法
func (y *YangGuo) NewKongfu() {
	fmt.Println(y.Name, "子类自己的新功夫！")
}

//子类重写父类的方法
/*func (y *YangGuo) ToadKongfu() {
	fmt.Println(y.Name, "的新蛤蟆功！")
}*/

func Test_Extends(t *testing.T) {
	o := &OuYangCrazy{Name: "欧阳疯", Age: 70} //创建父类
	o.ToadKongfu()                          //父类对象访问父类方法

	y := &YangGuo{OuYangCrazy{Name: "杨过", Age: 18}, "古墓"} //创建子类
	fmt.Println(y.Name)                                   //子类对象访问父类中有的字段
	fmt.Println(y.Address)                                //子类访问自己的字段

	y.ToadKongfu() //子类对象访问父类方法
	y.NewKongfu()  //子类访问自己的方法
	//y.ToadKongfu() //如果存在自己的方法 访问自己重写的方法
}
