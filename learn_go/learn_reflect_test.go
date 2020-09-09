package learn_go

import (
	"fmt"
	"reflect"
	"testing"
)

/**
通过反射获取类型对象与值对象:
	一个普通的变量包含两个信息，一个是类型type，一个是值value。
	type指的是系统中的原有的数据类型，如：int、string、bool、float32 等类型。
    在Go语言中可以通过 reflect.TypeOf() 函数获取任意值的类型对象，
    程序通过这个类型对象，可以获取任意的类型信息。
*/

func Test_reflect(t *testing.T) {
	a := 36
	fmt.Println("-------TypeOf------")
	aType := reflect.TypeOf(a) //通过反射获取变量a的type类型对象
	fmt.Println(aType.Name())  //获取类型名称为int

	fmt.Println("-------ValueOf------")
	aValue := reflect.ValueOf(a) //通过反射获取变量a的value类型对象
	fmt.Println(aValue.Int())    //获取具体的数值

	fmt.Println("-------Kind------")
	typeOfMyStruct := reflect.TypeOf(myStruct{Name: "结构体名称", Sex: 2})
	fmt.Println(typeOfMyStruct.Name()) //获取反射类型对象  myStruct
	fmt.Println(typeOfMyStruct.Kind()) //获取反射类型种类  struct

	fmt.Println("-------指针------")
	typeOfMyStruct2 := reflect.TypeOf(&myStruct{Name: "指针名称", Sex: 1})
	fmt.Println(typeOfMyStruct2.Elem().Name()) //获取指针类型指向的元素类型的名称
	fmt.Println(typeOfMyStruct2.Elem().Kind()) //获取指针类型指向的元素类型的种类

	fmt.Println("-------获取结构体成员字段的数量------")
	fieldNum := typeOfMyStruct.NumField() //获取结构体成员字段的数量
	for i := 0; i < fieldNum; i++ {
		fieldName := typeOfMyStruct.Field(i) //索引对应的字段信息
		fmt.Println(fieldName)
		name, err := typeOfMyStruct.FieldByName("Name") //根据指定的字符串返回对应的字段信息
		fmt.Println(name, err)
	}
}

type myStruct struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

//使用反射值对象获取任意值
func Test_Reflect_Value(t *testing.T) {
	a := 2020
	valueOf := reflect.ValueOf(a) //先通过reflect.ValueOf 获取反射的值对象
	fmt.Println(valueOf)
	//再通过值对象通过类型断言转换为指定类型
	fmt.Println(valueOf.Interface()) //转换为interface{} 类型
	fmt.Println(valueOf.Int())       //将值以int类型返回
}

//.Interface()将值以interface{}任意类型返回。
//还有各自对应的类型， .Int()、.Uint() 、.Floact() 、.Bool() 、.Bytes() 、.String()。

//通过反射获取结构体的成员字段的值
//reflect.Value 也提供了像获取成员类型的方法，用来获取成员的值。
func Test_Reflect_Struct_Value(t *testing.T) {
	h := haoJiaHuo{"好家伙", 20}
	fmt.Println(h)
	hOfValue := reflect.ValueOf(h)             //获取结构体的reflect.Value对象。
	for i := 0; i < hOfValue.NumField(); i++ { //循环结构体内字段的数量
		//获取结构体内索引为i的字段值
		fmt.Println(hOfValue.Field(i).Interface())
	}
	fmt.Println(hOfValue.Field(1).Type()) //获取结构体内索引为1的字段的类型

	fmt.Println("-------反射对象的空值处理------")
	var a *int                              //声明一个变量a为nil的空指针
	fmt.Println(reflect.ValueOf(a).IsNil()) //判断是否为nil 返回true

	//当reflect.Value不包含任何信息，值为nil的时候IsValid()就返回false
	fmt.Println(reflect.ValueOf(nil).IsValid())
}

type haoJiaHuo struct {
	Name string
	Age  int
}

//反射值对象reflect.Value提供了IsNil()方法判断空值。IsValid()方法判断是否有效。

//reflect.Value值对象支持修改反射出的元素的值。
func Test_Reflect_Update_Value(t *testing.T) {
	//声明变量a
	a := 100
	fmt.Printf("a的内存地址为：%p\n", &a)
	//获取变量a的反射类型reflect.Value 的地址
	rf := reflect.ValueOf(&a)
	fmt.Println("通过反射获取变量a的地址:", rf)

	//获取a的地址的值
	rval := rf.Elem()
	fmt.Println("反射a的值：", rval)

	//修改a的值
	rval.SetInt(200)
	fmt.Println("修改之后反射类型的值为：", rval.Int())

	//原始值也被修改
	fmt.Println("原始a的值也被修改为：", a)
}

//使用反射修改值的方法
//SetInt(x)   设置值为x, 类型必须是int类型。
//SetUint(x)  设置值为x, 类型必须是uint类型。
//SetFloat(x) 设置值为x, 类型必须是float32或者float64类型。
//SetBool(x)  设置值为x, 类型必须是bool类型。
//SetBytes(x) 设置值为x, 类型必须是[]Byte类型。
//SetString(x)设置值为x, 类型必须是string类型。
