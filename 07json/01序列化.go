package main

import (
	"encoding/json"
	"fmt"
)

/*
对结构体进行json序列化，属性一定要大写（大写才对外可见），否则json包无法访问
*/

type Person struct {
	Name    string
	Age     int
	Gender  bool
	RMB     float64
	Hobbies []string
}

func NewPerson(name string, age int, gender bool, rmb float64, hobbies []string) *Person {
	person := new(Person)
	person.Name = name
	person.Age = age
	person.Gender = gender
	person.RMB = rmb
	person.Hobbies = hobbies
	return person
}

/*
init函数会被系统优先调用，用于main之前的初始化工作
*/
var person *Person

func init() {
	// 返回的person是一个指针*Person
	person = NewPerson("张铁蛋", 20, true, 20.5, []string{"抽烟", "喝酒", "烫头"})
	// 输出指针的值
	fmt.Println(*person)
}

/*
序列化结构体
*/
func main011() {
	fmt.Println("main")
	bytes, err := json.Marshal(*person)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}
	fmt.Println(string(bytes))
}

/*
序列化map
*/
func main012() {

}

/*
序列化切片
*/
func main013() {

}
