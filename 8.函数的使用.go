// 函数的写法
// 1：有返回值的函数
// 2：无返回值的函数

package main

import "fmt"

var a, b = 1, 2

// 无参函数的使用
func swap() {
	a, b = b, a
}

// 有参函数,实现元素的交换
// 这个函数实现不了交换元素的,值传递不支持对象的交换
// 需要使用指针实现
func swap2(x, y int) {
	x, y = y, x
}

func swap3(x, y *int) {
	*x, *y = *y, *x
}

// 无返回值函数

// 有返回值的函数
// 实现一个获取最大值操作
func getMax(x int, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
func main() {
	fmt.Println("---------未调用函数时------------")
	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("---------调用函数时------------")
	swap()
	// 如果想使用别的输出格式,比如使用逗号进行分割
	// fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Println("a = ", a, "b = ", b)
	// 使用有参格式的交换方式
	fmt.Println("---------有参格式的交换-----------")
	swap2(a, b)
	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("---------指针实现交换-----------")
	swap3(&a, &b)
	fmt.Println("a = ", a, "b = ", b)
	// 获取最大值
	fmt.Println("---------函数返回值的使用-----------")
	var a, b = 10, 20
	var c int = getMax(a, b)
	fmt.Println("c = ", c)

	// 简单指针的使用
	// 简单的理解就是指针就像是固定的座位,解引用(*)相当于座位上的人
	var p *int
	p = &a
	// 打印地址
	// &a 是取a的地址,p这个指针指向就是a的地址
	fmt.Println(&a)
	fmt.Println(p)
	// 二级指针的使用
	var pp **int
	// pp指向的是p的地址
	pp = &p
	fmt.Println(pp)
}
