package learn_go

import (
	"fmt"
	"testing"
)

type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func Test_Person(t *testing.T) {
	//实例化后并使用结构体
	p := Person{} //使用简短声明方式，后面加上{}代表这是结构体
	p.age = 2     //给结构体内成员变量赋值
	p.address = "陕西"
	p.name = "好家伙"
	p.sex = "女"
	fmt.Println(p.age, p.address, p.name, p.sex) //使用点.来访问结构体内成员的变量的值。

	//直接给成员变量赋值
	p2 := Person{age: 2, address: "陕西", name: "老李头", sex: "女"}
	fmt.Println(p2.age, p2.address, p2.name, p2.sex)
}

func Test_New_Person(t *testing.T) {
	p2 := Person{age: 2, address: "陕西", name: "老李头", sex: "女"}

	//1 使用结构体指针
	var p *Person
	p = &p2 //将p2 的地址赋给p
	fmt.Println(p)
	p.name = "好家伙" //修改p的值
	fmt.Println(p)
	fmt.Println(p2) //p2的值也被修改了

	fmt.Println("------------------")
	//2 使用new 创建结构体指针
	pnew := new(Person)
	fmt.Println(pnew)
	pnew.address = "陕西"
	pnew.age = 23
	pnew.name = "李书记"
	pnew.sex = "男"
	fmt.Println(pnew)
}

//使用函数来实例化结构体
func Test_Pointer_Person(t *testing.T) {
	p := newPerson("好家伙", 18, "男")
	fmt.Println(p.name, p.age, p.sex)
}

//返回一个结构体的指针
func newPerson(name string, age int, sex string) *Person {
	return &Person{
		name: name,
		age:  age,
		sex:  sex,
	}
}

// 结构体内的每一个字段，都有自己相应的数据类型，
// 如果结构体被实例化后，字段的默认值就是该字段类型的零值，int就是0，string就是"",如果是指针类型，默认就是nil。

//结构体一
type Prescription struct {
	name     string
	unit     string
	additive Prescription2
}

//结构体二
type Prescription2 struct {
	name string
	unit string
}

//也可以嵌套结构体指针
type Prescription3 struct {
	name     string
	unit     string
	additive *Prescription2
}

func Test_Prescription(t *testing.T) {
	p := Prescription{}
	p.name = "鹤顶红"
	p.unit = "1.2kg"
	p.additive = Prescription2{
		name: "砒霜",
		unit: "0.5kg",
	}
	fmt.Println(p)

	//结构体初始化可以使用上面两种格式将字段名和对应的值写在括号内，使用(字段名:值,)的格式填充
	//第二种初始化的方式，定义好结构体之后使用重新赋值的方式:使用(变量.字段名=值)的格式

	//嵌套结构体指针
	pr := Prescription2{}
	pr.name = "鹤顶红升级版"
	pr.unit = "2.2kg"

	pre := Prescription3{}
	pre.name = "砒霜+"
	pre.unit = "1.2kg"
	pre.additive = &pr
	fmt.Println(pre)
}
