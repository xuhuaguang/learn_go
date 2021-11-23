package learn_go

import (
	"fmt"
	"testing"
	"time"
)

// 我们在聊天过程中，有两条通道，一条专门负责发送消息给对方，另一条通道专门负责接收消息。
// 虽然可以使用for循环来遍历每个通道的数据，达到同时接收到多个通道的数据，但是效率就比较差了。
// 在Go语言中提供select关键字，可以同事响应多个通道的操作。

func Test_Select(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case <-ch1: // 如果ch1成功读到数据，则进行该case处理语句
				fmt.Println("成功获取ch1的数据：", <-ch1)
			case ch2 <- 1: //如果成功向ch2写入数据，则进行该case处理语句
				fmt.Println("成功向通道ch2中写入数据")
			case <-time.After(time.Second * 2):
				//使用time.After 设置超时响应。如果迟迟接收不到以上的case就会响应超时。
				fmt.Println("超时!!", time.After(time.Second*2))
			}
		}
	}()

	for i := 0; i < 10; i++ {
		ch1 <- i
		fmt.Println("ch1写入数据：", i)
	}
	for i := 0; i < 10; i++ {
		str := <-ch2
		fmt.Println("获取到ch2的数据：", str)
	}

	fmt.Println("时间记录：", time.After(time.Second*2))
	// select 会一直等待等到某个 case 语句完成，
	//也就是等到成功从 ch1 或者 ch2 中读到数据。则 select 语句结束。
}

func TestSelect2(t *testing.T) {
	ch := make(chan int)
	c := 0
	stopCh := make(chan bool)

	go Chann(ch, stopCh)

	for {
		select {
		case c = <-ch:
			fmt.Println("Receive1", c)
			fmt.Println("channel")
		case s := <-ch:
			fmt.Println("Receive2", s)
		case _ = <-stopCh:
			goto end
		}
	}
end:
}

func Chann(ch chan int, stopCh chan bool) {
	i := 10
	for j := 0; j < 10; j++ {
		ch <- i
		time.Sleep(time.Second)
	}
	stopCh <- true
}
