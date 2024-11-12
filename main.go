// go语言打印输出输入
package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	// 输出的区别
	// Println在输出的结尾是带上空格的
	// Printf/Print在输出的结尾是不带空格的
	// 以下两句代码的输出为hello world!hello world!------》也就是说不会自动换行
	// 换行可以通过\n来完成,也可以使用Println来是现实输出换行
	fmt.Print("hello world!")
	fmt.Printf("hello world!")

}
