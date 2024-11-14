// map的实现是一个哈希表,另外map作为函数参数时,是引用传递,这个需要特别注意
// 在局部函数对map的修改是会影响全局的map

package main

import "fmt"

func main() {
	// 创建一个map
	var myMap map[string]string
	if myMap == nil {
		fmt.Println("map为空")
	} else {
		fmt.Println("map不为空")
	}
	// map[string]string(nil) myMap是nil
	fmt.Printf("%#v\n", myMap)
	// 在进行插入之前,需要开辟map的空间
	myMap = make(map[string]string, 10)
	// 对map的数据进行插入操作
	myMap["中国"] = "北京"
	myMap["美国"] = "华盛顿"
	// map[string]string{"中国":"北京", "美国":"华盛顿"}
	fmt.Printf("%#v\n", myMap)

	// 第二种创建方式,直接通过make创建并初始化大小
	myMap1 := make(map[int]int, 10)
	// 使用for循环进行大小填充
	for i := 0; i < 10; i++ {
		myMap1[i] = i
	}
	// 输出
	for key, value := range myMap1 {
		fmt.Printf("key:%v,value:%v\n", key, value)
	}

	// 第三种初始化的方式
	myMap3 := map[int]int{
		1: 1,
		2: 2,
	}
	fmt.Println("myMap3的大小len = ", len(myMap3))

	// map中存一个数组的形式
	myMap4 := make(map[int][]int)
	// 向map中插入元素
	// myMap4[1]是一个切片
	myMap4[1] = append(myMap4[1], 1)
	myMap4[2] = append(myMap4[2], 1)
	// map[int][]int{1:[]int{1}, 2:[]int{1}}
	fmt.Printf("%#v", myMap4)
}
