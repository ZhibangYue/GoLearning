package main

import "fmt"

/*
程序上线前必须扫灭、处理所有的恐慌，因为程序会在恐慌处终止运行
内建函数recover可以使崩溃的程序复活，并返回恐慌的内容
recover的返回值就是程序的死亡原因
*/

func main021() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("致死的原因是", err)
		} else {
			fmt.Println("程序正常结束")
		}
	}()
	fmt.Println("开始")
	panic("emmm")
	fmt.Println("更多功能")
}
