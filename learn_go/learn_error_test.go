package learn_go

import (
	"errors"
	"fmt"
	"testing"
)

// 通常情况一个函数如果有错误，都会在返回值最后一个，返回一个error类型的错误，
// 根据这个值来判断是否是非nil的值，如果是nil表示没有错误，如果nil不为空，则需要进行错误处理。

func Test_Error(t *testing.T) {
	err := errors.New("错误信息")
	fmt.Println(err)

	num, err2 := Calculation(0)
	fmt.Println(num, err2)
}

//通过内置errors包创建错误对象来返回
func Calculation(divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("错误:除数不能为零")
	}
	return 100 / divisor, nil
}
