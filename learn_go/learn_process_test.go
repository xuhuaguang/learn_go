package learn_go

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	if true {
		fmt.Println("进入判断数据")
	} else {
		fmt.Println("进入判断数据")
	}
}

func TestSwitch(t *testing.T) {
	num := 1
	switch num {
	case 1, 2, 3:
		fmt.Println("num符合其中某一个 执行代码")
		break
	case 4, 5, 6:
		fmt.Println("执行此代码")
		break
	}

	switch name := "hello"; name {
	case "hello":
		fmt.Println("hello")
	case "world":
		fmt.Println("world")
	}
}

func Test_Break_continue(t *testing.T) {
flag:
	for i := 1; i < 10; i++ {
		for j := 1; j < i; j++ {
			fmt.Println(i, j)
			if j == 5 {
				break flag
			}
		}
		fmt.Println(i)
	}
}

func Test_Goto(t *testing.T) {
TestLabel: //标签
	for a := 20; a < 35; a++ {
		if a == 25 {
			a += 1
			goto TestLabel
		}
		fmt.Println(a)
		a++
	}
}

//冒泡排序
func Test_Bubbling(t *testing.T) {
	array := []int{2, 5, 1, 7, 4, 9}
	BubblingASC(array)
	BubblingDESC(array)
}

//冒泡排序 正序，大的靠后 小的靠前。
func BubblingASC(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] > values[j] { //左右两边数据对比
				values[i], values[j] = values[j], values[i] //数据交换
			}
		}
	}
	fmt.Println(values)
}

//冒泡排序 倒序, 大的靠前 小的靠后。
func BubblingDESC(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[i] < values[j] { //左右两边数据对比
				values[i], values[j] = values[j], values[i] //数据交换
			}
		}
	}
	fmt.Println(values)
}
