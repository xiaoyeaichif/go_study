package main

import (
	"fmt"
	"reflect"
)

// 声明变量时,go的底层会有一个pair(类型,值)
// 也就是 pair(int,10)
func main() {
	var num float64 = 3.1415926
	fmt.Println("type =", reflect.TypeOf(num))
	fmt.Println("Value =", reflect.ValueOf(num))
}
