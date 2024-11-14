// 接口实现多态interface
// 实现多态的几个要求
// 1:需要有父类(接口),继承关系,子类重写父类的函数,父类的指针或者引用指向子类

package main

import "fmt"

// 定义一个父类接口类
// 要实现多态,子类必须重写父类的函数
type Animal interface {
	Sleep()
	GetColor() string // 获取颜色
	GetType() string  // 获取类型
}

// 定义一个Cat结构体,也就是子类
type Cat struct {
	Color string
}

// 设置子类的函数
func (cat Cat) Sleep() {
	fmt.Println("小猫睡觉!")
}

// 获取小猫的颜色
func (cat Cat) GetColor() string {
	return cat.Color
}

// 获取动物的类型
func (cat Cat) GetType() string {
	return "Cat!"
}

// 实现一个Dog类
type Dog struct {
	Color string
}

func (dog Dog) Sleep() {
	fmt.Println("小狗在睡觉!")
}

// Dog的成员函数的书写
func (dog Dog) GetColor() string {
	return dog.Color
}

// 获取动物的类型
func (dog Dog) GetType() string {
	return "Dog!"
}

// 显示不同动物的输出
// showInfo不属于任何函数
func ShowInfo(animal Animal) { // 这个函数就能实现,不同的对象,调用不同的函数
	animal.Sleep()
	fmt.Println("animal = ", animal.GetType())
	fmt.Println("color = ", animal.GetColor())
}

func main() {
	// var animal Animal // animal其实是一个父类指针
	cat := Cat{"black"}
	// animal = cat
	ShowInfo(cat) // 也可以直接传入cat这个参数,表示传入的是一个子类对象
	fmt.Println("---------------------")
	// 使用另外的Dog类来调用这个接口
	dog := Dog{"white"}
	ShowInfo(dog)
}

/*
	小猫睡觉!
	animal =  Cat!
	color =  black
	---------------------
	小狗在睡觉!
	animal =  Dog!
	color =  white
*/
