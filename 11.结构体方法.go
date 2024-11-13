// 定义结构体方法
package main

import "fmt"

// 大写代表公有的
type Person1 struct {
	Name string
	Age  int
	Sex  string
}

// 结构体方法的定义
func (p Person1) GetInfo() {
	// fmt.Printf("姓名:%s,年龄:%d,性别:%s\n", p.Name, p.Age, p.Sex)
	fmt.Println("姓名:", p.Name, "年龄:", p.Age, "性别:", p.Sex)
}

// 修改信息的成员函数
func (p *Person1) SetInfo(name string, age int, sex string) {
	p.Name = name
	p.Age = age
	p.Sex = sex
}
func main() {
	// 结构体变量信息的获取
	var p1 = Person1{
		Name: "yejie",
		Age:  22,
		Sex:  "男",
	}
	p1.GetInfo()
	// 修改信息,将原先的信息进行修改
	p1.SetInfo("yinyue", 24, "女")
	p1.GetInfo()
}
