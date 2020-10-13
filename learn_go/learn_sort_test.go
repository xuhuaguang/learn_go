package learn_go

import (
	"fmt"
	"sort"
	"testing"
)

//sort包提供了排序切片和用户自定义数据集的函数
//并且 sort 包已经帮我们把[]int,[]float64,[]string 三种类型都实现了该接口，我们可以方便的调用

var array = []int{3, 8, 2, 5, 1, 7, 9, 4}

func TestSortIntSlice(t *testing.T) {
	s := sort.IntSlice{3, 8, 2, 5, 1, 7, 9, 4}
	fmt.Println("len:", s.Len()) //数组大小

	s.Sort()
	fmt.Println("正向排序", s) //正向排序 [1 2 3 4 5 7 8 9]

	s.Swap(1, 0)
	fmt.Println("通过下标交换值", s) //[2 1 3 4 5 7 8 9]

	fmt.Println(s.Less(1, 0)) //下标1中的值是否大于下标0中的值
}

type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

func (s StuScores) Len() int {
	return len(s)
}

//Less(): 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score < s[j].score
}

func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestSort(t *testing.T) {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}
	//StuScores 已经实现了 sort.Interface 接口 , 所以可以调用 Sort 函数进行排序(升序，从小到大)
	sort.Sort(stus)
	// 判断是否已经排好顺序，将会打印 true
	fmt.Println("IS Sorted?\n\t", sort.IsSorted(stus))
	// 打印排序后的 stus 数据
	fmt.Println("Sorted:\n\t", stus)

	sort.Sort(sort.Reverse(stus))
	// 打印排序后的 stus 数据(降序，从大到小)
	fmt.Println("Sorted:\n\t", stus)
}

func TestIntArraySort(t *testing.T) {
	//var array = []int{3, 8, 2, 5, 1, 7, 9, 4}
	sort.Ints(array)
	fmt.Println(array) //将会输出[1 2 3 4 5 7 8 9]

	//如果要使用降序排序，用Reverse() 方法
	sort.Sort(sort.Reverse(sort.IntSlice(array)))
	fmt.Println(array) // 将会输出[6 5 4 3 2 1]

	//如果要查找整数 x 在切片 a 中的位置，相对于前面提到的 Search() 方法，sort包提供了 SearchInts():
	//注意，SearchInts() 的使用条件为：切片a已经升序排序
	s := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(sort.SearchInts(s, 2)) //必须升序排序，否侧返回0
}
