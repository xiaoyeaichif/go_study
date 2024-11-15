// channel的使用
// 1：可读可写管道,只读管道,只写管道
// 2:可以将管道看成是一个先进先出的队列,先写进管道的后输出

package main

import "fmt"

func main() {
	// 生成一个可读可写管道
	var num = make(chan int, 3) // 创建一个大小为3的管道
	num <- 1
	num <- 2
	num <- 3
	// 输出管道的大小
	fmt.Println("管道的大小size = ", len(num), ",容量为cap = ", cap(num))
	// 在规定的容量之外写数据,会导致死锁
	// // fatal error: all goroutines are asleep - deadlock!
	// num <- 4

	// 请注意,管道是不存在自动扩容机制的,如果初始化时规定了容量大下,那么管道的容量就不能发生变化
	fmt.Println("管道的大小size = ", len(num), ",容量为cap = ", cap(num))
	// 将管道当中的数据读取出来
	for i := 0; i < cap(num); i++ {
		a := <-num // 管道的输出
		fmt.Println("管道当中的数据为:", a)
	}

	// 创建一个只读管道
	// var num1 = make(<-chan int, 3)
	// num1 <- 1 // 这句代码表示向只读的管道中写入数据,这是不允许的,所以会报错

	// 创建一个只写管道
	// 只允许写,不允许读
	var num2 = make(chan<- int, 3)
	num2 <- 1
	num2 <- 2 // 写入数据1,2
	// 尝试读取数据是不允许的
	a := <-num2
	// invalid operation: cannot receive from send-only channel num2 (variable of type chan<- int)
	fmt.Println("尝试获取一个数据:", a)

}
