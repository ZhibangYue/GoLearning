package main

import (
	"fmt"
	"time"
)

func DoTask(no int) {
	for i := 1; i < 10; i++ {
		fmt.Println(no)
		time.Sleep(500 * time.Millisecond)
	}
}

/*
主协程迅速开辟十条并发，但每个协程执行任务需要5秒。任务未完成，主协程就已结束，因此任务终止程序结束。
主死从随
主协程一死，子协程都陪死。
*/
func main021() {
	for i := 1; i < 10; i++ {
		go DoTask(i)
	}
	fmt.Println("main over")
}

// 百万级并发
