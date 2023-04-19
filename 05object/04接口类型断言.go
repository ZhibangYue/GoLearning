package main

import "fmt"

/*
接口类型断言，即接口的类型判断
*/

type Cat struct {
	name string
}

func (c *Cat) SayHelloWorld() {
	fmt.Println("哼哼哼", c.name)
}

func (c *Cat) GoDie() {
	fmt.Println("啊死了", c.name)
}

func (c *Cat) Live(food string) (energy int) {
	energy = 100
	fmt.Printf("%s食用了%s，产生了%d卡路里", c.name, food, energy)
	return energy
}

func main041() {
	// 创建一个切片
	animals := make([]Animal, 0)
	animals = append(animals, &Pig{"猪"})
	animals = append(animals, &Cat{"猫"})
	// switch类型断言
	for _, animal := range animals {
		switch animal.(type) {
		case *Cat:
			fmt.Printf("是猫,%v", animal)
		case *Pig:
			fmt.Printf("是猪,%v\n", animal)

		}
	}
	// 另一种类型断言的写法
	// 如果不是对应的结构体类型，就会返回nil和false
	for _, animal := range animals {
		cat, ok := animal.(*Cat)
		fmt.Println(animal, cat, ok)
		if ok {
			cat.GoDie()
		}
	}
}
