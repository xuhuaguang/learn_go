package learn_go

import (
	"fmt"
	"testing"
)

// panic宕机 recover恢复 panic 表示恐慌。
// 当程序遇到一个异常时候可以强制执行让程序终止操作，同时需要引入defer函数类延时操作后面的函数，
// 在defer函数中通过recover来恢复正常代码的执行，因为defer是根据先入后出原则，所以先被defer的函数会放在最后执行，
// recover需要放在第一个被执行，当遇到panic时候恢复正常的代码逻辑。同时也可将错误信息通过recover获取panic传递的错误信息。

func Test_Recover(t *testing.T) {
	test1()
}

func test1() {
	defer func() {
		ms := recover()            //这里执行恢复操作
		fmt.Println(ms, "恢复执行了..") //恢复程序执行,且必须在defer函数中执行
	}()
	defer fmt.Println("第1个被defer执行")
	defer fmt.Println("第2个被defer执行")
	for i := 0; i <= 6; i++ {
		if i == 4 {
			panic("中断操作") //让程序进入恐慌 终端程序操作
		}
	}

	defer fmt.Println("第3个被defer执行") //恐慌之后的代码是不会被执行的
}
