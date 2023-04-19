package main

import "fmt"

/*
接口
接口定义了一组共性，这些共性体现为接口的抽象方法
抽象方法就是只有方法定义，没有方法实现的方法
接口里有且只有抽象方法
接口里可以有多种不同的具体子类实现
接口的作用是为子类提供统一的API
*/

type Animal interface {
	// 出生
	SayHelloWorld()
	// 死亡
	GoDie()
	// 生活
	Live(food string) (energy int)
}

type Pig struct {
	name string
}

func (p *Pig) SayHelloWorld() {
	fmt.Println("哼哼哼", p.name)
}

func (p *Pig) GoDie() {
	fmt.Println("啊死了", p.name)
}

func (p *Pig) Live(food string) (energy int) {
	energy = 100
	fmt.Printf("%s食用了%s，产生了%d卡路里", p.name, food, energy)
	return energy
}

func main031() {
	// 接口与实现
	//var animal Animal
	// 不关心是什么子类实现的接口，只关心接口的共性。
	//animal = &Pig{"猪"}
	//animal.SayHelloWorld()

	animals := make([]Animal, 0)
	animals = append(animals, &Pig{"猪"})
	for _, i := range animals {
		i.Live("饲料")
		i.GoDie()
	}
}
