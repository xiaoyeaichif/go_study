// go语言的数组使用,与C++vector有什么不同

package main

import "fmt"

// 主函数
func main() {
	// 数组的声明,使用短链接的形式声明
	ans := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(ans); i++ {
		//fmt.Println(ans[i])
		if i == len(ans)-1 {
			fmt.Print(ans[i])
		} else {
			fmt.Print(ans[i], " ")
		}
	}
	fmt.Print("\n")
	// 在for循环中声明变量时候,只能使用短链接的形式
	var nums = []int{1, 2, 3, 4, 5}
	// 输出每个元素
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 {
			fmt.Print(nums[i])
		} else {
			fmt.Print(nums[i], " ")
		}
	}
	// 手动输出换行操作
	fmt.Print("\n")
	// 实现一个冒泡排序
	var nums2 = []int{1, 5, 2, 4, 3}
	// 实现冒泡排序
	n := len(nums2)
	// 外层表示排序的次数
	for i := 0; i < n-1; i++ {
		// 内层表示排序的
		for j := 0; j < n-i-1; j++ {
			// 相邻的两个元素进行比较
			if nums2[j+1] <= nums2[j] {
				nums2[j+1], nums2[j] = nums2[j], nums2[j+1]
			}
		}
	}
	// 输出
	for i := 0; i < n; i++ {
		fmt.Print(nums2[i], " ")
	}
}
