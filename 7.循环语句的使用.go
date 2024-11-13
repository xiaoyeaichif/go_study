// 1：循环语句的使用,需要注意的是,go语言没有while循环这个选型,与C++不同
// 2：go语言的循环语句有for,range,switch,select2：
// 3：break,continue

package main

import "fmt"

// 简单循环实验
func main() {
	fmt.Println("----------------------")
	for i := 0; i < 5; i++ {
		//fmt.Println("i=", i) // 默认会输出一个空格,并且换行---》不用空格就是使用Print
		fmt.Printf("i=%d\n", i)
	}
	fmt.Println("----------------------")
	// 嵌套循环的使用
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}
	fmt.Println("----------------------")
	// 使用for循环模拟while循环,其实在用法上基本没什么区别,只是在写法上有区别
	var a = 5
	for a > 0 {
		fmt.Println("a=", a)
		a--
	}
	fmt.Println("----------------------")
	// 使用break语句
	var b = 5
	for b > 0 {
		if b == 1 {
			break
		}
		fmt.Println("b=", b)
		b--
	}
	fmt.Println("----------------------")
	// switch---case
	var c = 8
	switch c {
	case 1:
		fmt.Println("c=1")
		break
	case 2:
		fmt.Println("c=2")
		break
	case 3:
		fmt.Println("c=3")
		// 选中直接跳出
		break
	case 4:
		fmt.Println("c=4")
		// 选中直接跳出
		break
	case 5:
		fmt.Println("c=5")
		break
	default:
		fmt.Println("暂未找到满足题意得条件,请重试你的输入!")
	}
	// continue语句得使用
	for d := 0; d < 5; d++ {
		if d == 3 {
			continue
		}
		fmt.Println(d)
	}
}
