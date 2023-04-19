package main

import "fmt"

func main041() {
	fmt.Println("亲爱的患者，欢迎来到我院！")
	// 挂起一个延时任务，在当前函数返回（结束）前执行
	// 第一个defer会是最后一个执行，defer的执行顺序是倒序
	defer fmt.Println("再见，我院永远欢迎你！")
	defer fmt.Println("您的诊断已结束")
	fmt.Println("事务A")
	fmt.Println("事务B")
	fmt.Println("事务C")
}
