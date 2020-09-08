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

//拷贝
//浅拷贝指的是拷贝的引用地址，修改拷贝过后的数据,原有的数据也被修改
//深拷贝是指将值类型的数据进行拷贝的时候，拷贝的是数值本身，所以值类型的数据默认都是深拷贝
func Test_Copy(t *testing.T) {
	//使用range循环获取元素中的值进行拷贝
	slice := []int{1, 2, 3, 4}
	s2 := make([]int, 0)
	for _, v := range slice {
		s2 = append(s2, v)
	}
	fmt.Println(slice) //结果 [1 2 3 4]
	fmt.Println(s2)    //结果 [1 2 3 4]

	fmt.Println("-----------使用深拷贝数据函数: copy(目标切片,数据源)-----------")
	//copy(目标切片,数据源)  深拷贝数据函数
	s3 := []int{1, 2, 3, 4}
	s4 := []int{7, 8, 9}

	/*copy(s3, s4)    //将s4拷贝到s3中(替换同下标中的值)
	fmt.Println(s3) //结果 [7 8 9 4]
	fmt.Println(s4) //结果 [7 8 9]*/

	/*copy(s4, s3[2:]) //将s3中下标为2的位置 到结束的值 拷贝到s4中
	fmt.Println(s3)  //结果 [1 2 3 4]
	fmt.Println(s4)  //结果 [3 4 9]*/

	copy(s4, s3)    //将s3拷贝到s4中
	fmt.Println(s3) //结果 [1 2 3 4]
	fmt.Println(s4) //结果 [1 2 3]
}

//切片的删除
func Test_Slice_Delete(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice[2:]) //从数组下标2开始到结尾位置处的下标，包含下标2的数据  [3 4 5 6]
	fmt.Println(slice[:3]) //从数组下标3开始到开始位置处的下标，不包含下标3的数据   [1 2 3]
	fmt.Println(slice[:])  //全部数据 [1 2 3 4 5 6]

	fmt.Println("-----------删除切片中元素的方法-----------")
	//方法一 获取切片指定位置的值 重新赋值给当前切片
	method1Slice := []int{1, 2, 3, 4}
	method1Slice = method1Slice[1:] //删除切片中开头1个元素  结果 [2,3,4]

	//方法二 使用append不会改变当前切片的内存地址
	method1Slice = []int{1, 2, 3, 4}
	fmt.Println(method1Slice[:0])
	fmt.Println(method1Slice[1:])
	method1Slice = append(method1Slice[:0], method1Slice[1:]...) // 删除开头1个元素
	fmt.Println(method1Slice)
}

//删除指定的下标元素
func Test_Slice_Delete2(t *testing.T) {
	slice := []int{1, 2, 3, 4}
	i := 2                                    // 要删除的下标为2
	fmt.Println(slice[:i])                    //[1,2,3]
	fmt.Println(slice[i+1:])                  //[4]
	slice = append(slice[:i], slice[i+1:]...) // 删除中间1个元素
	fmt.Println(slice)                        //结果[1 2 4]

	fmt.Println("-----------删除切片结尾的方法-----------")
	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(slice2))
	slice2 = slice2[:len(slice2)-2] // 删除最后2个元素
	fmt.Println(slice2)             // 结果 [1,2]
}

func Test_Map(t *testing.T) {
	//1、声明map 默认值是nil
	//	var m1 map[string]int
	//2、使用make声明
	//  m2 := make(map[string]int)
	//3、直接声明并初始化赋值map的方法
	//	m3 := map[string]int{"语文": 89, "数学": 90}

	var m1 map[int]string         //只是声明 nil
	var m2 = make(map[int]string) //创建
	m3 := map[string]int{"语文": 89, "数学": 23, "英语": 90}

	fmt.Println(m1 == nil) //true
	fmt.Println(m2 == nil) //false
	fmt.Println(m3 == nil) //false

	//map 为nil的时候不能使用 所以使用之前先判断是否为nil
	if m1 == nil {
		m1 = make(map[int]string)
	}

	//1、存储键值对到map中  语法:map[key]=value
	m1[1] = "小猪"
	m1[2] = "小猫"

	//2、获取map中的键值对  语法:map[key]
	val := m1[2]
	fmt.Println(val)

	//3、判断key是否存在   语法：value,ok:=map[key]
	val, ok := m1[1]
	fmt.Println(val, ok) //结果返回两个值，一个是当前获取的key对应的val值。二是当前值否存在，会返回一个true或false。

	//4、修改map  如果不存在则添加， 如果存在直接修改原有数据。
	m1[1] = "小狗"

	//5、删除map中key对应的键值对数据 语法: delete(map, key)
	delete(m1, 1)

	//6、获取map中的总长度 len(map)
	fmt.Println(len(m1))

}

//遍历map
func Test_Map_Foreach(t *testing.T) {
	//map的遍历
	//因为map是无序的 如果需要获取map中所有的键值对
	//可以使用 for range
	m1 := make(map[int]string)
	m1[1] = "张无忌"
	m1[2] = "张三丰"
	m1[3] = "常遇春"
	m1[4] = "胡青牛"
	//遍历map
	for key, val := range m1 {
		fmt.Println(key, val)
	}

	fmt.Println("---------华丽的分割线----------")
	//map结合Slice
	//创建一个map存储第一个人的信息
	map1 := make(map[string]string)
	map1["name"] = "张无忌"
	map1["sex"] = "男"
	map1["age"] = "21"
	map1["address"] = "明教"

	//如果需要存储第二个人的信息则需要重新创建map
	map2 := make(map[string]string)
	map2["name"] = "周芷若"
	map2["sex"] = "女"
	map2["age"] = "22"
	map2["address"] = "峨眉山"

	//将map存入切片 slice中
	s1 := make([]map[string]string, 0, 2)
	s1 = append(s1, map1)
	s1 = append(s1, map2)
	//遍历map
	for key, val := range s1 {
		fmt.Println(key, val)
	}
}

