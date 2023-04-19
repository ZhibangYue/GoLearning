package main

import (
	"fmt"
	"strconv"
)

/*
类型转换
整型和浮点型可以直接强制转换
字符串可以通过strconv
*/
func main091() {
	var a int = 123
	var b float64 = 456.78
	aFloat := float64(a)
	fmt.Printf("aFloat的类型是%T,值是%v\n", aFloat, aFloat)
	bInt := int16(b)
	fmt.Printf("bInt的类型是%T，值是%v\n", bInt, bInt)
	resultInt, _ := strconv.ParseInt("123", 0, 64)
	fmt.Printf("result的类型是%T，值是%v\n", resultInt, resultInt)
	resultFloat, _ := strconv.ParseFloat("456.78", 64)
	fmt.Printf("result的类型是%T，值是%v\n", resultFloat, resultFloat)
}

func main092() {
	var a uint8 = 23
	// 0001,0111
	fmt.Println(a >> 3)
	fmt.Println(a << 3)
	// 左移5位，溢出
	fmt.Println(a << 5)
	var b int8 = 23
	// 0001,0111
	fmt.Println(b >> 3)
	// 左移3位，1011,1000，最高位为1，负数，减一取反得原码
	fmt.Println(b << 3)
	// 左移5位，溢出
	fmt.Println(b << 5)

}
