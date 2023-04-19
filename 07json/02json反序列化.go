package main

import (
	"encoding/json"
	"fmt"
)

var jsonStr = `{
	"name":"张铁蛋",
	"age": 20,
	"RMB": 20.5
}`
var jsonStr2 = `[
{
	"name":"张铁蛋",
	"age": 20,
	"RMB": 20.5
},
{
	"name":"张全蛋",
	"age": 20,
	"RMB": 0.5
}
]`

/*
将json反序列化为map
*/
func main021() {
	retMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &retMap)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(retMap)
}

/*
将json序列化为切片
*/
func main022() {
	retSlice := make([]map[string]interface{}, 0)
	err := json.Unmarshal([]byte(jsonStr2), &retSlice)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(retSlice)
}

type Human struct {
	Name string
	Age  int
	RMB  float64
}

/*
将json反序列化为结构体，结构体属性可以比json里更多，但不能更少
*/
func main023() {
	human := new(Person)
	// human := new(Human)
	err := json.Unmarshal([]byte(jsonStr), &human)
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	fmt.Println(human)
}
