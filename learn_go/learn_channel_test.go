package learn_go

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
)

// 当存在多个goroutine要传递某一个数据时，可以把这个数据封装成一个对象，
// 然后把对象的指针传入channel通道中，另一个goroutine 从通道中读取这个指针。
// 同一时间只允许一个goroutine访问channel通道里面的数据。所以go就是把数据放在了通道中来传递，而不是共享内存来传递。
// go启动协程的方式就是使用关键字go，后面一般接一个函数或者匿名函数

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

//缓冲通道
// 前面的channel通道都是非缓冲通道，每一次发送和接收都是阻塞式的。
// 一个发送操作，对应一个接收操作，如果发送后未接收，就是阻塞的。
// 同样对于接收者来说另一个发送之前它也是阻塞的。
// 缓冲通道指的是有一个缓冲区，对于发送数据是将数据发送到缓冲区。当缓冲区满了之后才会被阻塞。
func Test_Channel_Size(t *testing.T) {
	//定义一个缓冲区大小为5的通道
	ch1 := make(chan int, 5)
	ch1 <- 1 //向缓冲区放入数据1 因为缓冲区的大小为5 放入一个1之后 还有四个空的缓冲区  所以还未阻塞
	ch1 <- 2
	ch1 <- 3
	ch1 <- 4
	ch1 <- 5 //此时缓冲区已经满 如果再加入 则会进入阻塞状态
	//继续添加时会造成死锁 因为缓冲区满了 一直没有读取
	ch1 <- 6 //fatal error: all goroutines are asleep - deadlock!
	fmt.Println("main end")
}

// 对于缓冲区中的数据，是先进先出的原则。
// 第一个进去的第一个被读取到，可以理解为，一个从管道的一端放入，一个从管道的另一端读取。

//4、定向通道
//而定向通道表示： 要么是只读通道，要么是只写通道。
//chan <- T 只写通道
//<- chan T 只读通道

func Test_Channel_Direct(t *testing.T) {
	ch1 := make(chan int)   //双向通道
	ch2 := make(chan<- int) //只写通道
	ch3 := make(<-chan int) //只读通道

	//=========1================
	//如果创建时候创建的就是双向通道
	//则在子协程内部写入数据，读取的时候不受影响。
	go WriteOnly(ch1)
	data2 := <-ch1
	fmt.Println("获取到只写通道中的数据是", data2)

	//=========2================
	//如果将定向通道ch2只写通道，作为参数传递。
	//则不能读取到写回来的数据。
	go WriteOnly(ch2)
	//data := <-ch2 //不能读取会报错：invalid operation: <-ch2 (receive from send-only type chan<- int)

	go ReadOnly(ch1) //这里可以传ch1 双向通道
	ch1 <- 20        //向通道ch1中写入数据

	//=========3================
	go ReadOnly(ch3) //传递单向通道ch3 就无法向通道中写入数据

	fmt.Println("结束")
}

//只读
func ReadOnly(ch <-chan int) {
	data := <-ch
	fmt.Println("读取到通道的数据是：", data)
}

//只写
func WriteOnly(ch chan<- int) {
	//如果传进来的原本是双向通道
	//但是函数本身接收的是一个只写的通道，则在此函数内部只允许写入数据不允许读取数据
	//所以单向通道往往是作为参数传递
	ch <- 10
	fmt.Println("只写通道结束")
}

//5、死锁
// 死锁是指两个或两个以上的协程的执行过程中，由于竞争资源而阻塞的现象，如果没有外力介入,则无法继续进行下去。
// 死锁的出现的情况有很多种，但都大多数都是因为资源竞争和数据通信的时候引起的。

//创建一个同步等待组的对象
var wgp sync.WaitGroup

func Test_Dead_Lock(t *testing.T) {
	wgp.Add(4) //设置同步等待组的数量
	go Sale1()
	go Sale2()
	go Sale3()
	wgp.Wait() //主goroutine进入阻塞状态
	fmt.Println("main end...")
}

func Sale1() {
	fmt.Println("func1...")
	wgp.Done() //执行完成 同步等待数量减1
}
func Sale2() {
	defer wgp.Done()
	fmt.Println("func2...")
}
func Sale3() {
	defer wgp.Done() //使用延时执行来减去执行组的数量
	fmt.Println("func3...")
}

//一个通道在一个主goroutine协程里同时进行读和写。也会造成死锁。
func Test_Dead_Lock2(t *testing.T) {
	c := make(chan int)
	c <- 100 //向通道中写入数据
	a := <-c //读取通道中的数据
	fmt.Println(a)
}

//协程开启之前就放数据,还没有准备好，就放数据，就会造成死锁。
func Test_Dead_Lock3(t *testing.T) {
	c := make(chan int)
	c <- 88
	go func() {
		<-c
	}()
}

//如果你运行上面第一段代码，你会发现什么结果都没有.
//当你使用go启动协程之后,后面没有代码了，这时候主线程结束了，这个协程还没来得及执行就结束了
//https://zhuanlan.zhihu.com/p/74047342
func Test_XieCheng(t *testing.T) {
	go say("Hello World")
}

func say(s string) {
	println(s)
}

//上述Test_XieCheng解决方法
//简单说明一下用法，var是声明了一个全局变量wgWay1，
//类型是sync.WaitGroup，wgWay1.add(1) 是说我有1个协程需要执行，
//wgWay1.Done 相当于 wgWay1.Add(-1) 意思就是我这个协程执行完了。
//wgWay1.Wait() 就是告诉主线程要等一下，等协程都执行完再退出。
var wgWay1 = sync.WaitGroup{}

func Test_WaitGroupWay1(t *testing.T) {
	wgWay1.Add(1)
	go sayWay1("Hello World")
	wgWay1.Wait()
}

func Test_WaitGroupWay2(t *testing.T) {
	wgWay1.Add(5)
	for i := 0; i < 5; i++ {
		go sayWay1("Hello World: " + strconv.Itoa(i))
	}
	wgWay1.Wait()
}

func sayWay1(s string) {
	println(s)
	wgWay1.Done()
}
