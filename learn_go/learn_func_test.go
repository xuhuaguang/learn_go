package learn_go

import (
	"fmt"
	"testing"
)

func Test_Func(t *testing.T) {
	//声明一个变量f  类型是一个函数类型
	var f func()
	//将自定义函数myfunc 赋给变量f
	f = myfunc
	//调用变量f 相当于调用函数myfunc()
	f()
}

//自定义函数
func myfunc() {
	fmt.Println("myfunc...")
}

//匿名函数
func Test_Anonymous_Func(t *testing.T) {
	func() {
		fmt.Println("匿名函数")
	}()

	func(a int, b int) {
		fmt.Println(a, b)
	}(1, 2)

	res := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(res) //打印匿名函数返回值

	//匿名函数可以作为另一个函数的参数
	//匿名函数可以作为另一个函数的返回值
}

func Test_Oper(t *testing.T) {
	res2 := oper(20, 12, add)
	fmt.Println(res2)

	//匿名函数作为回调函数直接写入参数中
	res3 := oper(2, 4, func(a, b int) int {
		return a + b
	})
	fmt.Println(res3)
}

func add(a, b int) int {
	return a + b
}

func reduce(a, b int) int {
	return a - b
}

//oper就叫做高阶函数
//fun 函数作为参数传递则fun在这里叫做回调函数
func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b) //20 12 0x49a810A   第三个打印的是传入的函数体内存地址
	res := fun(a, b)  //fun 在这里作为回调函数 程序执行到此之后才完成调用
	return res
}

//defer函数调用时候，参数已经传递了，只不过代码暂时不执行而已。等待主函数执行结束后，才会去执行。
func Test_Defer(t *testing.T) {
	defer test(1) //第一个被defer的，函数后执行
	defer test(2) //第二个被defer的，函数先执行
	test(3)       //没有defer的函数，第一次执行

	//执行结果
	//3
	//2
	//1

	test2()
	//执行结果
	//2
	//2
	//1

}

func test(s int) {
	fmt.Println(s)
}

func test2() {
	a := 2
	test(a)
	defer test(a) //此时a已经作为2 传递出去了 只是没有执行
	a++           //a++ 不会影响defer函数延迟执行结果
	test(a)
}
