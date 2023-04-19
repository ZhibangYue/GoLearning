package main

import "fmt"

/*
变量作用域
全局变量：写在函数外，当前包下的所有函数都可以访问的变量（整个包）。
局部变量：定义在函数内部，其它函数无法访问（作用域是函数内部）。
函数参数和返回值都属于局部变量。

生命周期
全局变量生命周期与程序相同，程序结束时，全局变量才会释放。
局部变量生命周期与函数相同，函数执行结束，局部变量的内存就释放。
要避免无度定义全局变量。
*/

var publicBus = "aaa"

func main041() {
	fmt.Println(publicBus)
}
