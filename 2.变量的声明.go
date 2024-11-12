package main

import "fmt"

func main() {
	// 使用 var 关键字声明变量并赋值
	// var类似于C++的auto关键字,能做到自动推导的形式
	// 不同的点在于var会自动初始化变量,而auto是不会自动初始化的
	var name string = "Alice" // 显示类型推导,也可以写成var x = 42----->自动推导x为int
	var age int = 30
	var isStudent bool = true

	// 打印变量
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)

	// 一次初始化多个变量
	// 变量的类型可以不一致
	var (
		a int    = 1
		b string = "hello world!"
		c bool   = true
	)
	// 输出
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	// 使用自动推导,而不是显示声明
	var userName = "yejie"
	var userAge = 24
	var userHigh = 174

	// 输出
	fmt.Println("userName:", userName)
	fmt.Println("userAge:", userAge)
	fmt.Println("userHigh:", userHigh)
}
