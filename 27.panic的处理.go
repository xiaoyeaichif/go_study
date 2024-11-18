// panic的使用
package main

import "fmt"

func div(a, b int) {
	// 使用defer函数捕获这个异常
	// 结合recover函数
	// recover只在defer函数中有用
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()
	fmt.Println("执行除法操作:", a/b)
}

func main() {
	// panic: runtime error: integer divide by zero
	// 需要捕获这个异常
	// fmt.Println(div(10, 0))
	div(10, 0)
}
