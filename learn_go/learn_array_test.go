package learn_go

import (
	"fmt"
	"testing"
)

func Test_Array(t *testing.T) {
	array := []int{1, 4, 7, 3, 6}
	//range方式循环数组
	for k, v := range array {
		fmt.Println("k==>", k, "<===v===>", v)
	}

	for e := range array {
		fmt.Println("----------", e)
	}

	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

func Test_Slice(t *testing.T) {
	slice := make([]int, 3, 5) //长度为3 容量为5  容量如果省略 则默认与长度相等也为3
	fmt.Println(slice)
	fmt.Println(len(slice), cap(slice))
	fmt.Println("---------华丽的分割线----------")

	var s []int
	s = append(s, 1, 2, 3)
	fmt.Println(len(s), s)

	fmt.Println("---------使用make函数创建切片----------")
	s2 := make([]int, 0, 5)
	fmt.Println(s2)
	s2 = append(s2, 1, 2)
	fmt.Println(s2)
	//因为切片可以扩容  所以定义容量为5 但是可以加无数个数值
	s2 = append(s2, 3, 4, 5, 6, 7)
	fmt.Println(s2) // [1,2,3,4,5,6,7]

	fmt.Println("---------添加一组切片到另一切片中----------")
	s3 := make([]int, 0, 3)
	s3 = append(s3, s2...) //...表示将另一个切片数组完整加入到当前切片中
	fmt.Println(s3)
}

func Test_Append_Slice(t *testing.T) {
	s1 := make([]int, 0, 3)
	fmt.Printf("地址%p,长度%d,容量%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 1, 2)
	fmt.Printf("地址%p,长度%d,容量%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 3, 4, 5)
	fmt.Printf("地址%p,长度%d,容量%d\n", s1, len(s1), cap(s1))
	//	地址0xc00000e540,长度0,容量3
	//	地址0xc00000e540,长度2,容量3
	//	地址0xc00000c330,长度5,容量6
	//1、容量成倍数扩充：3-》6-》12-》24
	//2、如果添加的数据容量够用，则地址不变。如果实现了扩容，地址就会发生改变成新的地址，旧的地址自动销毁。
}

//值传递与引用传递
//值类型：int、float、string、bool、array、struct 值传递是传递的数值本身，不是内存地址，将数据备份一份传给其他地址，本身不影响，如果修改不会影响原有数据。
//引用类型: slice、pointer、map、chan 等都是引用类型。 因为引用传递存储的是内存地址，所以传递的时候则传递是内存地址，因此会出现多个变量引用同一个内存。
func Test_Value_Addr(t *testing.T) {
	fmt.Println("-----------数组为值传递类型-----------")
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println(arr1, arr2) //[1 2 3 4] [1 2 3 4]  输出结果 arr1与arr2相同

	fmt.Println("-----------修改arr1中下标为2的值200-----------")
	arr1[2] = 200           //修改arr1中下标为2的值
	fmt.Println(arr1, arr2) //[1 2 200 4] [1 2 3 4] 结果arr1中结果改变,arr2中不影响
	//说明只是将arr1中的值给了arr2, 修改arr1中的值后并不影响arr2的值

	fmt.Println("-----------切片是引用类型-----------")
	//定义一个切片 slice1
	slice1 := []int{1, 2, 3, 4}
	slice2 := slice1            //将slice1的地址引用到slice2
	fmt.Println(slice2, slice2) //[1 2 3 4] [1 2 3 4]   slice1输出结果 slice2输出指向slice1的结果，
	slice1[2] = 200             //修改slice1中下标为2的值
	fmt.Println(slice1, slice2) //[1 2 200 4] [1 2 200 4] 结果slice1中结果改变,因为修改的是同一份数据
	//说明只是将slice1中的值给了slice2 修改slice1中的值后引用地址用的是同一份 slice1 和slice2 同时修改

	fmt.Printf("%p,%p\n", slice1, slice2) //0xc000012520,0xc000012520
	//切片引用的底层数组是同一个 所以值为一个地址 是引用的底层数组的地址
	fmt.Printf("%p,%p\n", &slice1, &slice2) //0xc0000044a0,0xc0000044c0
	//切片本身的地址
}
