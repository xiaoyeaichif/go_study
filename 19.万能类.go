// interface{}万能类型
// 使用interface{}万能类型，可以存储任何类型的数据，但是不能使用它的方法

package main

import "fmt"

// 实现一个万能类型,可以接收所有类型的参数
func Niubi(arg interface{}) {
	fmt.Println("my Niubi Function is called...")
	fmt.Printf("arg = %v,arg type = %T\n", arg, arg)

	// 判断类型是不是我们想要的类型
	// 断言的使用
	// ok返回的是一个bool值，如果断言成功,就是true,反之则为false
	// value返回的是arg对应的值
	value, ok := arg.(string)
	if ok {
		fmt.Println("arg = ", value)
	} else {
		fmt.Println("arg is not string type!")
	}
}

// 主函数
func main() {
	// 声明一个整数
	fmt.Println("-----------整数---------")
	a := 100
	Niubi(a)
	// 声明一个字符串
	fmt.Println("---------字符串---------")
	b := "yejie"
	Niubi(b)
	// 声明一个浮点数
	fmt.Println("---------浮点数---------")
	c := 3.1415926
	Niubi(c)
	// 声明一个布尔值
	fmt.Println("---------布尔值---------")
	flag := false
	Niubi(flag)
}
