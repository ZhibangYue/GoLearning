package main

import "fmt"

/*
闭包函数：返回函数的函数
为不同的人保存不同的状态
闭包函数的好处：使用同一份内层函数代码，创建出任意多个不同的函数对象，这些对象各自的状态都被保存在函数闭包（外层函数）中，各行其道互不干扰
不使用闭包就需要多个全局变量来保存进度
*/

// var progress int
//func Study(name string, hours int) int {
//	fmt.Printf("%s学习了%d小时\n", name, hours)
//	progress += hours
//	return hours
//}

func GetStudy(name string) func(int) int {
	var progress int
	study := func(hours int) int {
		fmt.Printf("%s学习了%d小时\n", name, hours)
		progress += hours
		return progress
	}
	fmt.Printf("study的类型是%T\n", study)
	return study
}
func main061() {
	//progress := Study("黑旋风", 5)
	//fmt.Printf("李逵的学习进度是%d/10000\n", progress)
	studyFunc := GetStudy("武松")
	studyFunc(5)
	studyFunc(3)
	ws := studyFunc(2)
	fmt.Printf("武松的学习进度是%d/10000\n", ws)
	studyFunc2 := GetStudy("李逵")
	studyFunc2(5)
	lk := studyFunc2(2)
	fmt.Printf("李逵的学习进度是%d/10000\n", lk)
}
