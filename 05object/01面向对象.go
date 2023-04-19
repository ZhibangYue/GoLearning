package main

import "fmt"

/*
面向对象中类的概念，go中是通过结构体来实现的，对象是结构体/类的实例
面向函数的三大特性：封装、继承、多态
继承性实现了低成本地扩展代码
多态性实现了大规模代码的组成和调度
封装的意义在于整理和简化代码的结构，它也是继承和多态的基础
*/

type Dog struct {
	name   string
	age    int
	gender bool
}

// Bite 结构体方法的定义：前置主语指针类型
func (d *Dog) Bite(who string) {
	fmt.Printf("%s咬了%s\n", d.name, who)
}

func main011() {
	// 创建空白对象
	dog := Dog{}
	// 为对象的属性赋值
	dog.gender = true
	fmt.Println(dog.gender)
	dog.Bite("you")

	// 创建对象的同时赋值（可以选择性地赋值）
	dog2 := Dog{
		name:   "dog",
		age:    12,
		gender: false,
	}
	fmt.Println(dog2.name)
	dog4 := GetDog()
	fmt.Println(dog4.name)
}

// 封装
// 通过内建函数new获取对象指针
// 封装的GetDog这一函数是可以传参的（工厂模式）

func GetDog() *Dog {
	// dogPtr是Dog类型的指针
	dogPtr := new(Dog)
	// 与C++不同，通过指针访问属性和方法与通过对象访问是一致的，不需要箭头
	dogPtr.age = 2
	dogPtr.name = "小黑"
	return dogPtr
}
