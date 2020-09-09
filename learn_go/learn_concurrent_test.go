package learn_go

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"testing"
	"time"
)

// go语言的并发属于go语言中一大亮点，其他语言创建并发是通过线程，
// 而go语言则通过协程，协程是一个轻量级的线程。进程或者线程在一台电脑中最多不能超过一万个，
// 而协程可以在一台电脑中创建上百万个也不会影响到电脑资源。

//Goroutine是如何执行的
//	与函数不同的是goroutine调用之后会立即返回，不会等待goroutine的执行结果，所以goroutine不会接收返回值。
//  把封装main函数的goroutine叫做主goroutine，main函数作为主goroutine执行，
//  如果main函数中goroutine终止了，程序也将终止，其他的goroutine都不会再执行。
func Test_Goroutine(t *testing.T) {
	go testGo1()
	go testGo2()
	for i := 0; i <= 5; i++ {
		fmt.Println("main函数执行", i)
	}
	time.Sleep(3000 * time.Millisecond) //加上休眠让主程序休眠3秒钟。
	fmt.Println("main 函数结束")
}

func testGo1() {
	for i := 0; i <= 10; i++ {
		fmt.Println("111测试子goroutine1", i)
	}
}

func testGo2() {
	for i := 0; i <= 10; i++ {
		fmt.Println("222测试子goroutine2", i)
	}
}

// 由结果可以看出，当主函数main执行完成后，子goroutine执行了一次整个程序就执行结束了，main函数并不会等待子goroutine执行结束。
// 一个goroutine的执行速度是非常快的，并且是主goroutine和子goroutine进行资源竞争，谁抢到资源多，谁就先执行。main函数是不会让着子goroutine的。
// 我们可以在主goroutine中加上时间休眠，可以看每一个goroutine执行过程。

//2、使用匿名函数创建Goroutine
//使用匿名函数创建goroutine时候在匿名函数后加上(),直接调用。

func Test_Anonymous_Goroutine(t *testing.T) {
	go func() {
		fmt.Println("匿名函数创建goroutine执行")
	}()

	fmt.Println("主函数执行")
}

//3、runtime包
func Test_Runtime(t *testing.T) {
	//获取当前GOROOT目录
	fmt.Println("GOROOT:", runtime.GOROOT())
	//获取当前操作系统
	fmt.Println("操作系统:", runtime.GOOS)
	//获取当前逻辑CPU数量
	fmt.Println("逻辑CPU数量：", runtime.NumCPU())

	//设置最大的可同时使用的CPU核数  取逻辑cpu数量
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n) //一般在使用之前就将cpu数量设置好 所以最好放在init函数内执行

	//goexit 终止当前goroutine
	//创建一个goroutine
	go func() {
		fmt.Println("start...")
		runtime.Goexit() //终止当前goroutine
		fmt.Println("end...")
	}()
	time.Sleep(3 * time.Second) //主goroutine 休眠3秒 让子goroutine执行完
	fmt.Println("main_end...")
}

//4、Go语言临界资源安全
//什么是临界资源
//  指并发环境中多个协程之间的共享资源，如果对临界资源处理不当，往往会导致数据不一致的情况。
//  例如：多个goroutine在访问同一个数据资源的时候，其中一个修改了数据，另一个goroutine在使用的时候就不对了。

//定义全局变量 表示救济粮食总量
var food = 10

func Test_Dead_Source(t *testing.T) {
	//开启4个协程抢粮食
	go Relief("灾民好家伙1")
	go Relief("灾民好家伙2")
	go Relief("灾民老李头1")
	go Relief("灾民老李头2")

	//让程序休息5秒等待所有子协程执行完毕
	time.Sleep(5 * time.Second)
}

//定义一个发放的方法
func Relief(name string) {
	for {
		if food > 0 { //此时有可能第二个goroutine访问的时候 第一个goroutine还未执行完 所以条件也成立
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond) //随机休眠时间
			food--
			fmt.Println(name, "抢到救济粮 ，还剩下", food, "个")
		} else {
			fmt.Println(name, "别抢了 没有粮食了。")
			break
		}
	}
}

//执行结果：
//灾民老李头2 抢到救济粮 ，还剩下 9 个
//灾民好家伙1 抢到救济粮 ，还剩下 8 个
//灾民老李头2 抢到救济粮 ，还剩下 7 个
//灾民好家伙1 抢到救济粮 ，还剩下 6 个
//灾民老李头2 抢到救济粮 ，还剩下 5 个
//灾民老李头1 抢到救济粮 ，还剩下 4 个
//灾民好家伙2 抢到救济粮 ，还剩下 3 个
//灾民好家伙1 抢到救济粮 ，还剩下 2 个
//灾民老李头2 抢到救济粮 ，还剩下 1 个
//灾民老李头1 抢到救济粮 ，还剩下 0 个
//灾民老李头1 别抢了 没有粮食了。
//灾民老李头2 抢到救济粮 ，还剩下 -1 个
//灾民老李头2 别抢了 没有粮食了。
//灾民好家伙1 抢到救济粮 ，还剩下 -2 个
//灾民好家伙1 别抢了 没有粮食了。
//灾民好家伙2 抢到救济粮 ，还剩下 -3 个
//灾民好家伙2 别抢了 没有粮食了。

// 以上代码出现负数的情况，也是因为Go语言的并发走的太快了，当有一个协程进入执行的时候还没来得及取出数据，另外一个协程也进来了，
// 所以会出现负数的情况，那么如何解决这样的问题，我们不能用休眠的方法让程序等待，
// 因为你并不知道程序会多久执行结束，到底应该让程序休眠多长时间。

//5、sync同步包
// sync同步包，是Go语言提供的内置同步操作，保证数据统一的一些方法，
// WaitGroup 等待一个goroutine的集合执行完成，也叫同步等待组。
// 使用Add()方法，来设置要等待一组goroutine 要执行的数量。
// 用Done() 方法来减去执行goroutine集合的数量。
// 使用Wait()方法让主goroutine也就是main函数进入阻塞状态，等待其他的子goroutine执行结束后，main函数才会解除阻塞状态。

//创建一个同步等待组的对象
var wg sync.WaitGroup

func Test_WaitGroup(t *testing.T) {
	wg.Add(3) //设置同步等待组的数量
	go Relief1()
	go Relief2()
	go Relief3()
	wg.Wait() //主goroutine进入阻塞状态
	fmt.Println("main end...")
}

func Relief1() {
	fmt.Println("func1...")
	wg.Done() //执行完成 同步等待数量减1
}
func Relief2() {
	defer wg.Done()
	fmt.Println("func2...")
}
func Relief3() {
	defer wg.Done() //推荐使用延时执行的方法来减去执行组的数量
	fmt.Println("func3...")
}

//6、互斥锁
// 互斥锁，当一个goroutine获得锁之后其他的就只能等待当前goroutine执行完成之后解锁后才能访问资源。
// 对应的方法有上锁Lock()和解锁Unlock()。

//创建一把锁
var mutex sync.Mutex

func Test_Lock(t *testing.T) {
	wg.Add(4)
	//开启4个协程抢粮食
	go Relief4("灾民好家伙")
	go Relief4("灾民好家伙2")
	go Relief4("灾民老李头")
	go Relief4("灾民老李头2")
	wg.Wait() //阻塞主协程，等待子协程执行结束
}

//定义一个发放的方法
func Relief4(name string) {
	defer wg.Done()
	for {
		//上锁
		mutex.Lock()
		if food > 0 { //加锁控制之后每次只允许一个协程进来，就会避免争抢
			food--
			fmt.Println(name, "抢到救济粮 ，还剩下", food, "个")
		} else {
			mutex.Unlock() //条件不满足也需要解锁 否则就会造成死锁其他不能执行
			fmt.Println(name, "别抢了 没有粮食了。")
			break
		}
		//执行结束解锁，让其他协程也能够进来执行
		mutex.Unlock()
	}
}

//创建一把读写锁 可以是他的指针类型
var rwmutex sync.RWMutex

//7、读写锁
func Test_Read_Write(t *testing.T) {
	wg.Add(3)
	go ReadTest(1)
	go WriteTest(2)
	go ReadTest(3)
	wg.Wait()
	fmt.Println("======main结束======")
}

//读取数据的方法
func ReadTest(i int) {
	defer wg.Done()
	fmt.Println("======准备读取数据======")
	rwmutex.RLock() //读上锁
	fmt.Println("======正在读取...", i)
	rwmutex.RUnlock() //读取操作解锁
	fmt.Println("======读取结束======")
}

func WriteTest(i int) {
	defer wg.Done()
	fmt.Println("======开始读写数据======")
	rwmutex.Lock() //写操作上锁
	fmt.Println("======正写数据...", i)
	time.Sleep(1 * time.Second)
	rwmutex.Unlock() //写操作解锁
	fmt.Println("======写操作结束======")
}
