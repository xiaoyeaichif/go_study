package main

import "fmt"

// 定义函数对象
// 类似于C++中使用using一个函数对象的形式
// 定义一个函数对象,能接收任意个整形参数,并且返回值是一个整数
type func1 func(int, ...int) int

// 函数对象必须要有载体
func add(a int, b ...int) int {
	sum := 0
	for _, v := range b {
		sum += v
	}
	fmt.Printf("%v----%T\n", b, b)
	return sum
}

// 主函数的调用
func main() {
	fmt.Println("---------函数对象的调用---------")
	// 常见函数对象
	// 声明一个函数对象
	var f1 func1
	fmt.Printf("%v----%T\n", f1, f1)
	f1 = add
	// fmt.Println("%v\n", &f1)
	fmt.Printf("%v----%T\n", f1, f1)
	fmt.Println(f1(1, 2, 3, 1, 2, 3))
}
