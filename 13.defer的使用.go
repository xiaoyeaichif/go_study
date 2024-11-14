// defer的使用
// 1.defer语句会延迟到函数返回之前执行-------》有点像C++的析构函数-----》在函数结束之后才执行
// 比如说右return的存在,defer语句的执行也是在return之后的,和析构函数很像
// 另外需要记住的是defer是压栈的形式,也就是说defer语句会按照先进后出的顺序执行

package main

import "fmt"

func returnDefer() int {
	fmt.Println("return Defer")
	return 0
}

// return 函数的调用
func returnFunc() int {
	fmt.Println("return Func")
	return 0
}

// 执行两个函数
// returnFunc的执行要优先于returnDefer
func test() int {
	defer returnDefer()
	return returnFunc()
}

func main() {
	// defer语句的输出执行顺序是defer2----->defer1
	fmt.Println("---------defer2和defer1的执行顺序-----------")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")

	fmt.Println("---------defer和return的执行顺序-----------")
	// return Func
	// return Defer
	// 也就是说defer函数的执行是在最后,也就是程序执行完毕(和析构函数相同)
	test()

}

// 整体的执行顺序是主函数的当中的defer是最后执行的,所以这段代码的指向顺序为
/*
test--->return Func
test--->return Defer
main--->defer2
main--->defer1
*/
