// 封装的使用以及成员函数的定义

package main

import "fmt"

// 类的首字母大写,代表是公有的类
// 成员函数,成员变量首字母大写代表公有的,小写代表私有的

// 下面的类是一个公有的类,但是成员函数和成员变量都是私有的
// 相当于类
type Student struct {
	name string
	age  int
}

// 成员函数,获取姓名
func (s Student) getName() string {
	return s.name
}

// 成员函数,获取年龄
func (s Student) getAge() int {
	return s.age
}

// 设置姓名和年龄
// 成员函数的设置
func (s *Student) setName(name string, age int) {
	s.name = name
	s.age = age
}
func main() {
	// 初始化类
	var s1 Student
	s1.name = "yejie"
	s1.age = 25
	fmt.Println(s1)
	// 获取输出的值
	str := s1.getName()
	fmt.Println("name =", str)
	// 获取年龄
	age := s1.getAge()
	fmt.Println("age =", age)
	// 修改姓名和年龄输出
	s1.setName("yinyue", 20)
	// 输出姓名和年龄
	fmt.Println("name = ", s1.getName(), ", age = ", s1.getAge())
}
