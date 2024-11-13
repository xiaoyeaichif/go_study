// go中不存在class,都是使用struct构造结构体
// 进行实例化的时候,不管是用指针还是值初始化都可以用.号来进行，而不是C++中指针用->符号
// go中的type很像C++中的using
package main

import "fmt"

// 自定义类型
// 使用type关键字实现自定义类型
type myInt int // myInt是int的别名
/*
*** 相当于C++中的结构体
struct MyInt {
    int value;
};
*/

// type myFunc func(int, int) int

// 定义类型别名
type myFloat = float64 // 类似于C++中的    using MyFloat = float;

// 结构体的定义
// Person是结构体类型，首字母大写,表示公有的,如果首字母小写那就是私有的
type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var a myInt = 10
	fmt.Println("a = ", a)
	fmt.Printf("%v %T\n", a, a)
	// 适用类型别名
	var b myFloat = 3.14
	fmt.Printf("%v %T\n", b, b)

	// 使用结构体,结构体初始化
	var p1 Person // 实例化
	p1.name = "小明"
	p1.age = 18
	p1.sex = "男"
	// 值:{小明 18 男},类型:main.Person
	fmt.Printf("值:%v,类型:%T\n", p1, p1)
	// 打印详细的信息
	// 值:main.Person{name:"小明", age:18, sex:"男"},类型:main.Person
	fmt.Printf("值:%#v,类型:%T\n", p1, p1)

	// 使用指针初始化结构体
	var p2 = &Person{}
	p2.name = "小红"
	p2.age = 19
	p2.sex = "女"
	// 值:&{小红 19 女},类型:*main.Person
	fmt.Printf("值:%v,类型:%T\n", p2, p2)
	// 打印详细的信息
	// 值:&main.Person{name:"小红", age:19, sex:"女"},类型:*main.Person
	fmt.Printf("值:%#v,类型:%T\n", p2, p2)

	// 使用new关键字创建
	var p3 = new(Person)
	p3.name = "yejie"
	p3.age = 22
	p3.sex = "男"
	// 值:&{yejie 22 男},类型:*main.Person
	fmt.Printf("值:%v,类型:%T\n", p3, p3)
	// 打印详细的信息
	// 值:&main.Person{name:"yejie", age:22, sex:"男"},类型:*main.Person
	fmt.Printf("值:%#v,类型:%T\n", p3, p3)

	// 还可以多个变量一起初始化
	var p4 = Person{
		name: "yejie",
		age:  23,
		sex:  "男",
	}
	fmt.Printf("值:%v,类型:%T\n", p4, p4)
	fmt.Printf("值:%#v,类型:%T\n", p4, p4)

	// 结构体初始化一定要使用{},而不能使用()
	var p5 = Person{
		"yejie",
		23,
		"男",
	}
	fmt.Printf("值:%v,类型:%T\n", p5, p5)
	fmt.Printf("值:%#v,类型:%T\n", p5, p5)
}
