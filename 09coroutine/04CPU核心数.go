package main

import (
	"fmt"
	"runtime"
)

func main041() {
	fmt.Println("当前协程的可用CPU数为", runtime.NumCPU())
	// 将可用CPU逻辑核心数设置为4，并返回先前的配置
	gomaxprocs := runtime.GOMAXPROCS(4)
	fmt.Println(gomaxprocs) // 8
}
