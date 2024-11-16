// channel的使用
// 1：可读可写管道,只读管道,只写管道
// 2:可以将管道看成是一个先进先出的队列,先写进管道的后输出

/*

无缓冲管道的写操作：
	当一个goroutine尝试向无缓冲管道写入数据时，如果此时没有其他goroutine从该管道读取数据，写操作会阻塞。
	一旦有其他goroutine从管道中读取数据，写操作会被解除阻塞，数据被成功写入管道。
无缓冲管道的读操作：
	当一个goroutine尝试从无缓冲管道读取数据时，如果此时没有其他goroutine向该管道写入数据，读操作会阻塞。
	一旦有其他goroutine向管道中写入数据，读操作会被解除阻塞，数据被成功读取。
*/

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
	// var num2 = make(chan<- int, 3)
	// num2 <- 1
	// num2 <- 2 // 写入数据1,2
	// // 尝试读取数据是不允许的
	// // a := <-num2
	// // invalid operation: cannot receive from send-only channel num2 (variable of type chan<- int)
	// fmt.Println("尝试获取一个数据:", a)

	// 使用匿名函数表示来操作channel
	var num3 = make(chan int)
	// num3 <- 10
	go func() {
		// 延迟输出
		defer fmt.Println("匿名函数go程执行完毕.....")
		// 输出匿名函数正在运行
		fmt.Println("子go程正在运行....")

		// 向管道channel写入数据
		num3 <- 100
	}()

	// 读取匿名函数管道中的数据
	// 如果当前num3这个channel缓冲区的大小为0,主协程阻塞在这个写操作上,
	/*
		子go程正在运行....
		匿名函数go程执行完毕.....
		从匿名函数管道中读取的数据为: 100
		---------主协程执行完毕-----------
	*/
	temp := <-num3
	fmt.Println("从匿名函数管道中读取的数据为:", temp)
	fmt.Println("---------主协程执行完毕-----------")

}
