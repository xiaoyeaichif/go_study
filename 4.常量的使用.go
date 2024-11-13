// 常量的定义：常量必须使用const来声明
// 常量的声明只存在于布尔型,数字型,字符串这三种类型中

// 2：常量的声明与var声明(短变量声明有什么区别)？
// ---常量是使用const修饰的,声明之后,常量值不可以进行修改
// ---var声明相当于普通的定义,是可以再进行修改的,另外,短变量声明其实就是var的一种特殊格式
// var a = 10-----> var a int = 10------>a:=10

// 3:iota的使用----》计数器
/*
*1：
*
 */

package main

import "fmt"

func main() {
	const a int = 10
	fmt.Println("a=", a)
	// 常量是不可以修改的
	// a = 100;
	// 声明bool常量
	const flag bool = false // 显示定义上类型,更容易理解
	fmt.Println("flag = ", flag)
	// 声明字符串常量
	const str string = "yejie"
	fmt.Println("str = ", str)

	// iota的使用
	const new_a = iota
	fmt.Println("new_a =", new_a)
	// 多个声明,默认会加1
	const (
		new_b = iota
		new_c
		new_d
	)
	// 输出
	fmt.Println("new_b =", new_b)
	fmt.Println("new_c =", new_c)
	fmt.Println("new_d =", new_d)
}
