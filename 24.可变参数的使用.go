// go可变参数的使用

package main

import "fmt"

// [1 2 3 4 5 6 7 8 9]----[]int ---->也就是说可变参数最后是以切片的形式出现的
func sumFun(x ...int) int {
	fmt.Printf("%x----%T\n", x, x)
	// 输出所有的参数之和
	sum := 0
	for key, value := range x {
		fmt.Println(key, value)
		sum += value
	}
	return sum
}

// 可变参数的另外一种形式
func sumFun2(x int, y ...int) int {
	fmt.Printf("%x----%T\n", y, y)
	sum := x
	for _, value := range y {
		sum += value
	}
	return sum
}

func main() {
	sumFun(1, 2, 3, 4, 5, 6, 7, 8, 9)
	ans := sumFun(1, 2, 3, 4, 5)
	fmt.Println(ans)
	// 分割线
	fmt.Println("--------可变参数的另外一种形式----------")
	ans2 := sumFun2(1, 2, 3, 4, 5)
	fmt.Println(ans2)
}
