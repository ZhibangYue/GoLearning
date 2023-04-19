package main

import "fmt"

/*
继承
*/

type PoliceDog struct {
	Dog
	skill string
}

func (pd *PoliceDog) DoPoliceJob() {
	fmt.Printf("%s正在执行%s工作\n", pd.name, pd.skill)
}

func main021() {
	pd := PoliceDog{}
	// 父类属性
	pd.name = "警犬"
	pd.skill = "emm"
	fmt.Println(pd.name)
	pd.DoPoliceJob()
	// 通过子类对象访问父类方法和成员
	pd.Bite("aaa")

	pd2 := PoliceDog{Dog{name: "战狼"}, "emmmm"}
	// 创建时可以部分赋值 pd3 := PoliceDog{Dog: Dog{}}
	pd2.DoPoliceJob()
	pd3 := GetPoliceDog("警犬2", 3, true, "eat")
	pd3.DoPoliceJob()
}

// GetPoliceDog 工厂模式创建警犬对象
func GetPoliceDog(name string, age int, gender bool, skill string) *PoliceDog {
	pd := new(PoliceDog)
	pd.gender = gender
	pd.name = name
	pd.age = age
	pd.skill = skill
	return pd
}
