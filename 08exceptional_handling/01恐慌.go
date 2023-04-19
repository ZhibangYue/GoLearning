package main

import "fmt"

/*
代码在运行时如果出现异常，系统会报出恐慌（panic）并终止运行
IDE和终端打印的恐慌日志，包含了恐慌信息及其所在代码行
代码在交付使用前要经过充分测试，处理掉一切可能的恐慌
*/

/*
下标越界
*/
func main011() {
	var mySlice = []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice[5])
}

/*
报出恐慌
内建函数panic可以实现自己报出恐慌
目的是预测程序运行时可能出现的异常情形，并提示当前代码的调用者错误信息
*/
func main012() {
	fmt.Println(123)
	panic("恐慌的原因")
	fmt.Println(456)
}
