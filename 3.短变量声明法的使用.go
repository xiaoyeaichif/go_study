// 短变量声明法的使用,用处较多
package main

import "fmt"

// 需要注意短变量声明法不能用在全局声明,只能用在函数的内部,也就是局部变量的使用上
// 全局变量的使用还是要使用var来进行
// ans := 1 //----->这就是错误的做法

// 匿名变量
func getUserInfo() (string, int) {
	return "yejie", 18
}

func main() {

	// 短变量声明法
	//fmt.Println(ans)

	// 变量的声明
	userName := "yejie"
	userAge := 18
	// 输出
	fmt.Println("userName:", userName)
	fmt.Println("age:", userAge)
	// 打印变量的格式
	/*
	******* 常见的格式有哪些
	* %T------》输出变量的类型
	* %v------》输出变量的值
	* %t------》输出布尔值
	* %s------》输出字符串
	* %d------》输出整数
	* %f------》输出浮点数
	* %.2f  ----》输出浮点数,保留两位小数
	* %5d 设置输出宽度为 5，不足部分用空格填充
	 */
	fmt.Printf("userName type:%T,userAge type:%T\n", userName, userAge)

	// 匿名变量
	var name, age = getUserInfo()
	fmt.Println("name:", name, "age:", age)
	// 匿名变量的声明
	var new_name, _ = getUserInfo()
	fmt.Println("new_name:", new_name)
	// 重新赋值操作,不使用短变量声明

	name, age = getUserInfo()
	fmt.Println("name:", name, "age:", age)

	// 使用短变量声明法,接受到函数的返回值,必须要有一个新声明的变量
	name, newAge := getUserInfo()
	fmt.Println("name:", name, "newAge:", newAge)
}
