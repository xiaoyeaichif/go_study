// 1：切片的定义---->本质上就是一个动态数组,底层有扩容机制,和vector很像
// 2：切片与数组的区别
/*
** 1：切片是引用类型，数组是值类型,也就说在作为形参时,对切片进行修改,就是对原值进行修改
** 2：切片是引用类型,数组是值类型,可以减少拷贝的消耗
** 3：切片的大小是可以动态变化的,数组的值初始化之后,不可以再次进行修改
 */

// 3：切片具备的方法
/*
	1:在末尾增加一个元素append
	2:截取切片s[:]
	3:排序函数的使用sort
	4:获取切片的大小len
	5:获取切片的容量cap
	6:切片的复制copy
*/

// 切片扩容的规则
/*
1：初始容量为 0：
	第一次追加元素时，容量从 0 扩容到 1。

2：容量小于 1024：
	容量在 1 到 1024 之间时，每次扩容都会翻倍。
例如，从 1 到 2，从 2 到 4，从 4 到 8，从 8 到 16。

3：容量大于或等于 1024：
	容量达到 1024 时，扩容后的容量会增加当前容量的一半。
例如，从 1024 到 1536，从 1536 到 2304
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 普通数组的定义,大小固定就是普通数组
	// 输出普通数组
	fmt.Println("--------普通数组的使用-------")
	var nums1 = [5]int{1, 2, 3, 4, 5} // 普通数组的定义
	nums2 := [5]int{5, 4, 3, 2, 1}    // 短变量声明
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println("--------切片的使用-------")
	// 大小不固定就是切片
	// 初始化的格式
	/*
		var myArray []int
		myArray = make([]int, 5)
	*/
	myArray := []int{1, 2, 3}
	// 输出切片的大小
	fmt.Println("myArray切片的长度len = ", len(myArray), ", myArray切片的容量cap = ", cap(myArray))

	// 不同的初始化格式
	var myArray1 []int
	// 指定初始化的len和cap,初始化长度为5,容量为10
	myArray1 = make([]int, 5, 10)
	fmt.Println("myArray1切片的长度len = ", len(myArray1), ", myArray1切片的容量cap = ", cap(myArray1))

	// 换一种初始化的方式----->初始化数组的长度为2,默认的容量也为2
	myArray2 := make([]int, 2)
	fmt.Println("myArray2切片的长度len = ", len(myArray2), ", myArray2切片的容量cap = ", cap(myArray2))

	// append函数的实验,以及扩容的实验
	myArray3 := make([]int, 1, 2)
	// 填充数组
	for i := 0; i < 1; i++ {
		myArray3[i] = i
	}
	myArray3 = append(myArray3, 1)
	fmt.Println("myArray3切片的长度len = ", len(myArray3), ", myArray3切片的容量cap = ", cap(myArray3))
	// 再次增加元素,使得myArray3进行扩容
	myArray3 = append(myArray3, 2)
	fmt.Println("myArray3切片的长度len = ", len(myArray3), ", myArray3切片的容量cap = ", cap(myArray3))
	// 扩容函数的实现,在1024个范围大小是双倍扩容
	// 大于1024是取原先容量的一半大小再进行扩容
	// 指定数组的容量
	myArray4 := make([]int, 1024, 1024)
	for i := 0; i < 1024; i++ {
		myArray4[i] = i
	}
	fmt.Println("myArray4切片的长度len = ", len(myArray4), ", myArray4切片的容量cap = ", cap(myArray4))
	// 往后加入一个元素之后,检查容量的变化
	myArray4 = append(myArray4, 1024)
	fmt.Println("myArray4切片的长度len = ", len(myArray4), ", myArray4切片的容量cap = ", cap(myArray4))

	// 截取的使用
	myArray5 := make([]int, 10) // 大小为10,数组中所有元素都为0
	for i := 0; i < 10; i++ {
		myArray5[i] = i
	}
	fmt.Println("myArray5切片的长度len = ", len(myArray5), ", myArray5切片的容量cap = ", cap(myArray5))
	fmt.Println("---------截取的使用----------")
	// 截取
	// 截取的是左闭右开的区间[start,end)
	newArray5 := myArray5[1:5]
	fmt.Println(newArray5)
	// 需要注意上面的方式是用应用的方式截取的,修改其中一个都会对切片有影响的
	// 下面的输出明显就是修改了原切片,浅拷贝
	// [0 50 2 3 4 5 6 7 8 9]
	// [50 2 3 4]
	newArray5[0] = 50
	fmt.Println(myArray5)
	fmt.Println(newArray5)

	// 不修改原切片，执行深拷贝
	newArray6 := make([]int, len(myArray5))
	// 执行拷贝动作
	copy(newArray6, myArray5)
	fmt.Println(newArray6)
	newArray6[0] = 1
	// 检查拷贝结果
	// 执行的是深拷贝
	fmt.Println(myArray5)
	fmt.Println(newArray6)

	// 切片的排序使用
	myArray7 := []int{3, 1, 2, 5, 4, 7, 8, 0}
	// 排序前的输出
	fmt.Println("---------排序前的输出---------")
	for i := 0; i < len(myArray7); i++ {
		fmt.Print(myArray7[i], " ")
	}
	fmt.Print("\n")
	// 排序后的输出
	fmt.Println("---------排序后的输出---------")
	sort.Ints(myArray7) // 整数型排序
	for i := 0; i < len(myArray7); i++ {
		fmt.Print(myArray7[i], " ")
	}
}
