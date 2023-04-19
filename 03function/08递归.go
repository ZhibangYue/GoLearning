package main

import "fmt"

/*
循环的效率远高于递归
*/

func GetLoop(n int) int {
	if n == 1 {
		return 1
	} else {
		return n + GetLoop(n-1)
	}
}

// GetFibonacci 递归求斐波那契数列
func GetFibonacci(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return GetFibonacci(n-1) + GetFibonacci(n-2)
	}
}

// GetFibonacci2 循环求斐波那契数列
func GetFibonacci2(n int) int {
	var a, b = 1, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

func main081() {
	fmt.Println(GetLoop(5))
	for i := 0; i < 10; i++ {
		fmt.Println(GetFibonacci(i))
		fmt.Println(GetFibonacci2(i))
	}

}
