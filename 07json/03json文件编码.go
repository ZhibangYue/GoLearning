package main

import (
	"encoding/json"
	"os"
)

type Person2 struct {
	Name    string
	Age     int
	Gender  bool
	RMB     float64
	Hobbies []string
}

func NewPerson2(name string, age int, gender bool, rmb float64, hobbies []string) *Person {
	person := new(Person)
	person.Name = name
	person.Age = age
	person.Gender = gender
	person.RMB = rmb
	person.Hobbies = hobbies
	return person
}

func main031() {
	person2 := NewPerson("张铁蛋", 20, true, 30.5, []string{"抽烟"})
	dstFile, _ := os.OpenFile("D:\\go_craft\\GoLearning\\07json\\json文件.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 066)
	defer dstFile.Close()
	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(person2)
	if err != nil {
		return
	}

}

/*
切片编码进文件
*/
func main032() {
	person2 := NewPerson("张铁蛋", 20, true, 30.5, []string{"抽烟"})
	person3 := NewPerson("张全蛋", 20, true, 30.5, []string{"抽烟"})
	human := make([]Person, 0)
	// person2 和 person3 是指针，把它们的值添加进来
	human = append(human, *person2, *person3)
	dstFile, _ := os.OpenFile("D:\\go_craft\\GoLearning\\07json\\json文件.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 066)
	defer dstFile.Close()
	encoder := json.NewEncoder(dstFile)
	err := encoder.Encode(human)
	if err != nil {
		return
	}
}
