package main

import (
	"fmt"
	"strings"
)

/*
分割与拼接
separator 分隔符
Split(s, sep string) 分割
SplitAfter(s, sep string) 保留分隔符的分割
SplitN(s, sep string, n int) 指定分成n段
SplitAfterN(s, sep string,n int) 保留分隔符的分割，但是指定分成n段
Join([]string, sep string)拼接
*/
func main051() {
	fmt.Println(strings.Split("他好 我也好 大家好才是真的好", " "))
	fmt.Println(strings.Split("他好,我也好,大家好才是真的好", ","))
	fmt.Println(strings.SplitAfter("他好,我也好,大家好才是真的好", ","))
	fmt.Println(strings.SplitN("他好,我也好,大家好才是真的好,好好学习", ",", 3))
	fmt.Println(strings.Join([]string{"男", "女", "abcd"}, "*"))
}
