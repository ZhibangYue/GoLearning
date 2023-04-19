package main

import (
	"fmt"
	"runtime"
	"time"
)

func main031() {
	for i := 1; i <= 2; i++ {
		go func() {
			for j := 1; j <= 10; j++ {
				fmt.Printf("协程%d:%d\n", i, j)
			}
		}()
	}
	time.Sleep(1 * time.Second)
}

/*
输出中子协程中的i都是3
因为协程需要一定的初始化时间，当子协程真正执行时，主协程for循环已结束，i为3。
子协程初始化完时，主协程已执行一段时间了。
是由主协程开辟子协程
*/

// 解决方案，在创建子协程的时候把值定下来传给他
func main032() {
	for i := 1; i <= 2; i++ {
		go func(no int) {
			for j := 1; j <= 10; j++ {
				fmt.Printf("协程%d:%d\n", no, j)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
}

/*
runtime.Gosched()出让当前协程的资源，调低当前协程的优先级，使该协程最后进行
*/
func main033() {
	for i := 1; i <= 5; i++ {
		go func(no int) {
			if no == 3 {
				runtime.Gosched()
			}
			for j := 1; j <= 10; j++ {
				fmt.Printf("协程%d:%d\n", no, j)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
}
