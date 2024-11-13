// 运算符 + - * / %
// 运算符的使用上和C++编程语言没什么不同,所以需要注意的点会少一点

package main

import (
	"fmt"
)

func main() {
	// 加法
	var a int = 10
	var b int = 20
	var c int = a + b
	fmt.Println("c =", c)
	// 乘法
	d := a * b
	fmt.Println("d =", d)
}
