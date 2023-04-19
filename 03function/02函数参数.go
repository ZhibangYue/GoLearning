package main

import "fmt"

// Add0 无参函数:参数列表是空的
func Add0() {
	fmt.Println("add")
}

// Add2 有参函数
func Add2(a int, b int) {
	sum := a + b
	fmt.Println("参数之和是", sum)
}

// Add4 多个参数，不同类型
func Add4(a, b int, c float64, d float64) {
	aFloat := float64(a)
	bFloat := float64(b)
	sum := aFloat + bFloat + c + d
	fmt.Println("参数之和是", sum)
}

// Add5 不定长参数
func Add5(args ...int) {
	fmt.Printf("args的类型是%T\n", args)
	fmt.Printf("args的长度是%d\n", len(args))
	fmt.Println(args)
}

/*
Add6 不定长参数与定长参数混用
不定长必须写在最后一个
*/
func Add6(a, b int, args ...float64) {
	aFloat := float64(a)
	bFloat := float64(b)
	sum := aFloat + bFloat
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	fmt.Println("参数的和是", sum)
}

func main021() {
	//Add0()
	//Add2(1, 2)
	//Add4(1, 3, 2.5, 6.7)
	Add5(1, 2, 3, 4)
	Add6(1, 2, 3.4, 5.6, 7)
}
