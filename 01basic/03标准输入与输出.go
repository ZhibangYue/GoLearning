package main

import (
	"fmt"
	"strconv"
)

// 每次接受一个用户输入
func main05() {
	var a, b string
	fmt.Println("请输入相加的两个数：")
	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println("a+b", a+b)

	aInt, _ := strconv.ParseInt(a, 0, 0)
	bInt, _ := strconv.ParseInt(b, 0, 0)
	fmt.Println(aInt + bInt)
}

func main06() {
	fmt.Println("请输入两个名字")
	var p1, p2 string
	fmt.Scan(&p1, &p2)
	fmt.Println("No.1:", p1)
	fmt.Println("No.2:", p2)
}
