package main

import "fmt"

// 简单判断语句的使用
// 注意的点else必须接在if判断结束的}后面,不然会报错
// 嵌套判断语句的使用
func main() {
	a := 10
	b := 20
	if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a <= b")
	}
	// 嵌套判断语句的使用
	if a > b {
		fmt.Println("a > b")
	} else {
		if a == b {
			fmt.Println("a == b")
		} else {
			fmt.Println("a < b")
		}
	}
	// 结合if-else if-else语句使用
	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a == b")
	}
}
