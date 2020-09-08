package learn_go

import (
	"fmt"
	"testing"
)

//Go语言中通过&获取变量的地址。
//通过*获取指针所对应的变量存储的数值。

func Test_Pointer(t *testing.T) {
	//定义一个变量
	a := 2
	fmt.Printf("变量A的地址为%p", &a) //通过%p占位符, &符号获取变量的内存地址。
	//变量A的地址为0xc000072090

	//创建一个指针
	//指针的声明 通过 *T 表示T类型的指针
	var i *int     //int类型的指针
	var f *float64 //float64类型的指针
	fmt.Println(i) // < nil >空指针
	fmt.Println(f)

	//因为指针存储的变量的地址 所以指针存储值
	i = &a
	fmt.Println(i)  //i存储a的内存地址0xc000072090
	fmt.Println(*i) //i存储这个指针存储的变量的数值2
	*i = 100
	fmt.Println(*i) //100
	fmt.Println(a)  //100通过指针操作 直接操作的是指针所对应的数值
}

//指针的指针，也就是存储的不是具体的数值了，而是另一个指针的地址。
func pointer() {
	a := 2
	var i *int      //声明一个int类型的指针
	fmt.Println(&a) //0xc00000c1c8
	i = &a          //将a的地址取出来放到i里面
	fmt.Println(&i) //0xc000006028
	var a2 **int    //声明一个指针类型的指针
	a2 = &i         //再把i的地址放进a2里面
	fmt.Println(a2) //获取的是a2所对应的数值0xc000006028也就是i的地址
}

//数组指针
func Test_Array_Pointer(t *testing.T) {
	//创建一个普通的数组
	arr := [3]int{1, 2, 3}
	fmt.Println(arr)

	//创建一个指针 用来存储数组的地址 即：数组指针
	var p *[3]int
	p = &arr       //将数组arr的地址，存储到数组指针p上。
	fmt.Println(p) //数组的指针 &[1 2 3] 后面跟数组的内容

	//获取数组指针中的具体数据 和数组指针自己的地址
	fmt.Println(*p) //指针所对应的数组的值
	fmt.Println(&p) //该指针自己的地址0xc000006030

	//修改数组指针中的数据
	(*p)[0] = 200
	fmt.Println(arr) //修改数组中下标为0的值为200   结果为：[200 2 3]

	//简化写法
	p[1] = 210       //意义同上修改下标为1的数据
	fmt.Println(arr) //结果： [200 210 3]
}

//指针数组
//其实就是一个普通数组，只是存储数据类型是指针。
func Test_Pointer_Array(t *testing.T) {
	//定义四个变量
	a, b, c, d := 1, 2, 3, 4

	arr1 := [4]int{a, b, c, d}
	arr2 := [4]*int{&a, &b, &c, &d} //将所有变量的指针，放进arr2里面

	fmt.Println(arr1) //结果为：[1 2 3 4]
	fmt.Println(arr2) //结果为：[0xc00000c1c8 0xc00000c1e0 0xc00000c1e8 0xc00000c1f0]

	arr1[0] = 100                //修改arr1中的值
	fmt.Println("arr1的值：", arr1) //修改后的结果为：[100 2 3 4]

	fmt.Println("a=", a) //变量a的值还是1，相当于值传递，只修改了数值的副本。

	//修改指针数组
	*arr2[0] = 200 //修改指针的值
	fmt.Println(arr2)
	fmt.Println("a=", a) //200  引用传递 修改的是内存地址所对应的值 所以a也修改了

	//循环数组，用*取数组中的所有值。
	for i := 0; i < len(arr2); i++ {
		fmt.Println(*arr2[i])
	}
}

//指针函数
//如果一个函数返回结果是一个指针，那么这个函数就是一个指针函数。
func Test_Pointer_Func(t *testing.T) {
	//函数默认为指针 只是不需要用 *
	a := fun1
	fmt.Printf("a的类型：%T,a的地址是%p 数值为%p \n", a, &a, a) //0x49c670 函数默认为指针类型
	a1 := fun1()
	fmt.Printf("a1的类型：%T,a1的地址是%p 数值为%v \n", a1, &a1, a1) //[]int,a1的地址是0xc0000044c0 数值为[1 2 3]

	a2 := fun2()
	fmt.Printf("a2的类型：%T,a1的地址是%p 数值为%v \n", a2, &a2, a2) //*[]int,a1的地址是0xc000006030 数值为&[1 2 3 4]
	fmt.Printf("a2的值为：%p\n", a2)                          //0xc000004520 指针函数返回的就是指针
}

//一般函数
func fun1() []int {
	c := []int{1, 2, 3}
	return c
}

//指针函数 返回指针
func fun2() *[]int {
	c := []int{1, 2, 3, 4}
	fmt.Printf("c的地址为%p：\n", &c) //0xc000004520
	return &c
}

func Test_Pointer_Func2(t *testing.T) {
	s := 10
	fmt.Println(s) //调用函数之前数值是10
	fun3(&s)
	fmt.Println(s) //调用函数之后再访问则被修改成2
}

//接收一个int类型的指针作为参数
func fun3(a *int) {
	*a = 2
}
