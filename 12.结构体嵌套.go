// 结构体嵌套

package main

import "fmt"

type PerSon struct {
	Name string
	Age  int
	Address
}

// 嵌套地址信息
type Address struct {
	Street string
	City   string
}

// PerSon类的成员函数设置方式
// 注意修改信息一定要传指针,传值的话只会处理值的拷贝,这个拷贝是不会影响原来的变量的
// 如果是获取信息,建议使用传值
func (p *PerSon) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func main() {
	// 结构体的初始化操作
	var p1 PerSon
	p1.Name = "yejie"
	p1.Age = 25
	p1.Address.City = "杭州"
	p1.Address.Street = "白杨街道"
	// 输出的详细信息
	// {yejie 25 {白杨街道 杭州}}
	fmt.Printf("%v\n", p1)
	// 输出详细的结构
	// main.PerSon{Name:"yejie", Age:25, Address:main.Address{Street:"白杨街道", City:"杭州"}}
	fmt.Printf("%#v\n", p1)

	// 修改嵌套当中的元素
	p1.Address.Street = "2号大街"
	// 输出详细的结构
	// main.PerSon{Name:"yejie", Age:25, Address:main.Address{Street:"2号大街", City:"杭州"}}
	fmt.Printf("%#v\n", p1)

	// 使用成员函数修改信息
	p1.SetInfo("yinyue", 20)
	// 输出修改后的信息------>个人信息有做修改
	// main.PerSon{Name:"yinyue", Age:20, Address:main.Address{Street:"2号大街", City:"杭州"}}
	fmt.Printf("%#v\n", p1)
}
