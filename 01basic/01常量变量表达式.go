// 包名 main包下的main函数是程序的入口
package main

//导入sdk（software developing kit）中的fmt包
import "fmt"

/*
一次性定义多个常/变量
定义在函数以外的对所有函数可见
首字母大写的成员对其他包可见
*/
const (
	lightSpeed = 300000
	months     = 12
)

var (
	yearSeconds = 365 * 24 * 3600
	days        = 30
)

func main01() {
	// 定义圆周率常量
	const pi = 3.14
	// 定义圆的半径
	var radius = 10.0
	// 定义圆的面积
	var area = 0.0
	// 使用表达式求圆的面积，赋值给面积变量
	area = pi * radius * radius

	fmt.Println(area)
}

// 访问全局常量和变量，写在函数外部的成员可以被当前包下的所有函数访问
func main02() {
	// 计算一光年的距离
	// var lightYearDistance = lightSpeed * yearSeconds
	// := 的意思是变量的声明赋值二合一，只能在函数内部使用
	lightYearDistance := lightSpeed * yearSeconds
	fmt.Println("一光年是", lightYearDistance, "公里")
}

func main03() {
	const a, b, c = 1, 2, 3
	var d, e, f = 4, 5, 6
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
