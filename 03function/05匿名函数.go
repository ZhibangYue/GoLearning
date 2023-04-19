package main

import (
	"fmt"
	"time"
)

// 匿名函数存在的原因：不考虑复用，只是为使代码块成为一个整体，用于封装一次性的代码

// 延时执行一个匿名函数
func main051() {
	fmt.Println("打开网络")
	fmt.Println("打开数据库")
	fmt.Println("打开文件")
	//defer fmt.Println("关闭网络")
	//defer fmt.Println("关闭数据库")
	//defer fmt.Println("关闭文件")
	// 只使用一次，无须复用，封装一次性的代码
	defer func() {
		fmt.Println("关闭网络")
		fmt.Println("关闭数据库")
		fmt.Println("关闭文件")
	}()

	fmt.Println("愉快地上网")
	fmt.Println("愉快地读写数据库")
	fmt.Println("愉快地读写文件")
}

// 并发执行一个匿名函数
func main052() {
	// 并发一个匿名任务
	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println("bbb")
			time.Sleep(500 * time.Millisecond)
		}
	}(107)
	for i := 0; i < 10; i++ {
		fmt.Println("aaa")
		time.Sleep(500 * time.Millisecond)
	}
}
