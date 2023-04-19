package main

import (
	"fmt"
	"time"
)

func TimeIt(f func(int) int, n int) {
	startTime := time.Now().UnixNano()
	f(n)
	endTime := time.Now().UnixNano()
	fmt.Printf("耗时%d纳秒\n", endTime-startTime)
}

func main091() {
	GetFibonacci0 := GetFibonacci2
	TimeIt(GetFibonacci0, 20)
}
