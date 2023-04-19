package main

import (
	"fmt"
	"time"
)

/*
for的死循环
*/
func main021() {
	for {
		fmt.Println("循环ing")
		time.Sleep(1 * time.Second)
	}
}

/*
for的有限循环
for 起始条件;循环条件;增长条件
*/
func main022() {
	var i int
	for i = 1; i < 100; i++ {
		fmt.Printf("i的值是%v", i)
	}
}

/*
打印2 4 6 8 10 …… 100
*/
func main023() {
	for i := 2; i <= 100; i += 2 {
		fmt.Println(i)
	}
}

/*
循环的嵌套
*/
func main024() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j, "*", i, "=", j*i, "\t")
		}
		fmt.Println()
	}
}
