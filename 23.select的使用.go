// go的select,与linx的select类似,但是go的select是多路监控多个channel,而不是I/O事件监听
// select选择的规则是基于通道的可用性

/*
	select 的选择规则
1：检查所有 case：
	select 会依次检查每个 case，看通道操作（发送或接收）是否可以立即进行。
	如果有多个 case 可以执行（例如多个通道都准备好了进行通信），select 会随机选择其中的一个 case 执行。这种随机性可以帮助避免选择偏向。
	如果只有一个 case 可以执行，那么 select 就会执行这个 case。
	如果所有的 case 都不可以立即执行（即所有通道都处于阻塞状态），select 语句会阻塞，直到其中一个 case 可以执行。
2：default 子句（可选）：
	如果 select 包含一个 default 子句，并且所有其他 case 都不能执行，那么会立即执行 default 子句，而不会阻塞。
	default 子句可以用来实现非阻塞的通道操作。
*/

/*
	弊端：select会依次遍历所有的case,确定哪些case是可以执行的,也就是说会遍历所有的case,然后再根据所有的
	case在其中选择一个,如果case特别多,会有性能消耗,此时应该优化channel的容量,让select可以快速判断出哪一个case可以执行
*/

package main

import "fmt"

// 斐波那契数列的监控
func Faboniqie(in, out chan int) {
	x, y := 1, 1
	for {
		select {
		case in <- x: // 检测到当前有协程阻塞在读事件,所以这里执行写操作
			x = y
			y = x + y
		case <-out: // 检测到当前有协程阻塞在写事件,所以这里执行读操作
			fmt.Println("quit.....!")
			return // 表示读事件完成,直接退出
		}
	}
}

// 尝试监控多个channel事件
func main() {
	// 生成两个读写事件channel
	var in, out = make(chan int), make(chan int)

	// 启动go程
	go func() {
		for i := 0; i < 6; i++ {
			// 从in管道中读取数据------》输出
			a := <-in // 当前子协程中的in还暂时没有数据,所以会阻塞在这个代码上,等待in中有数据往下执行
			fmt.Println(a)
		}
		// 往out中写入数据------》输入----》写事件
		out <- 0
	}()

	Faboniqie(in, out)
	// main go程的输出
	fmt.Println("main go exit ...")

	var temp = 'a'
	fmt.Printf("%v,%T\n", temp, temp)

}
