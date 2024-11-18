package main

import "fmt"

// 选择排序
// 找到局部最小值,每次都处理一个局部最小值
func SelectSort(arr []int) {
	// 获取切片的长度
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			// 找到全局最小值,然后进行交换即可
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 与最小值所在的index交换位置
		temp := arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
}

func main() {
	arr := []int{1, 4, 2, 3, 5, 6, 7, 1, 2, 3, 4}
	//
	fmt.Println("-------------排序前--------------")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%x ", arr[i])
	}
	fmt.Println("\n")
	SelectSort(arr)
	fmt.Println("-------------排序后--------------")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

}
