package learn_go

import (
	"fmt"
	"testing"
)

// 当存在多个goroutine要传递某一个数据时，可以把这个数据封装成一个对象，
// 然后把对象的指针传入channel通道中，另一个goroutine 从通道中读取这个指针。
// 同一时间只允许一个goroutine访问channel通道里面的数据。所以go就是把数据放在了通道中来传递，而不是共享内存来传递。

func Test_Channel(t *testing.T) {
	var channle chan int
	fmt.Printf("通道的数据类型:%T,通道的值:%v\n", channle, channle) //

	if channle == nil {
		channle = make(chan int)
		fmt.Printf("通过make创建的通道数据类型:%T,通道的值:%v,\n", channle, channle)
		//make创建后 通道的值为 0xc00005c060 也就是一个内存地址
		//所以channel 是一个引用类型的数据
	}
}

func Test_Channel_Use(t *testing.T) {
	ch1 := make(chan int)

	go func() {
		fmt.Println("======子协程执行======")
		data := <-ch1 //从通道中读取数据
		fmt.Println("读取到通道中的数据是:", data)
	}()

	ch1 <- 10 //往通道里放数据
	fmt.Println("======主协程结束======")
}

func Test_Channel_Close(t *testing.T) {
	ch1 := make(chan int)

	go func() {
		fmt.Println("======子协程执行======")
		for i := 0; i < 10; i++ {
			ch1 <- i //往通道中放数据
		}
		close(ch1) //结束发送数据  通知对方通道已经关闭了
	}()

	//通过for range循环读取通道中的数据,当通道关闭,循环也就结束了
	for v := range ch1 {
		fmt.Println("读取到的通道的数据：", v)
	}

	fmt.Println("主协程结束")
}
