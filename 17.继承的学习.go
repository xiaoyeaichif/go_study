// 使用结构体的嵌套来实现继承,不区分继承的隔离级别,如C++的public,protected,private
// 继承特性

package main

import "fmt"

// 生成一个地址类---->父类
type Country struct {
	City   string
	Street string
}

// 输出城市和街道
func (c Country) Show() {
	fmt.Println("城市为:", c.City, ",街道为:", c.Street)
}

// 生成一个国家类
type Address1 struct {
	Name string
	Country
}

// 输出一个国家的城市和街道以及具体名字
func (addre Address1) Show() {
	fmt.Println("城市为:", addre.City, ",街道为:", addre.Street, ", 国家为:", addre.Name)
}

func main() {
	// 输出一个城市名和街道名
	var c1 Country
	c1.City = "hangz"
	c1.Street = "2 hao street!"
	// 输出
	// fmt.Println("城市为:", c1.City, ",街道为:", c1.Street)
	c1.Show()

	// 使用继承(由结构体嵌套实现)
	var addre Address1
	addre.Name = "china!"
	addre.City = "hangz!"
	addre.Street = "2 hao street!"
	// 输出
	// fmt.Println("城市为:", addre.City, ",街道为:", addre.Street, ",名字为:", addre.Name)
	addre.Show()
}
