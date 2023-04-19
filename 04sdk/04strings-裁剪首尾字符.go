package main

import (
	"fmt"
	"strings"
)

/*
空白字符包括空格、制表符\t，换行符\n，回车符\r。
TrimSpace(s string) 返回将s首尾空白去除的副本
TrimPrefix(s, prefix string) 返回将s首部prefix去除的副本
TrimSuffix(s, suffix string) 返回将s尾部suffix去除的副本
TrimFunc(s string, func(r rune) bool) 将s的首尾字符代入匿名函数func，如返回true则去除该字符，否则保留，返回处理后的字符串
*/
func main041() {
	fmt.Println(strings.TrimSpace(" h   e l  l  o  "))
	fmt.Println(strings.TrimSpace("\n h   e l  l  o  \n"))
	fmt.Println(strings.TrimPrefix("a:hello b", "a:"))
	fmt.Println(strings.TrimSuffix("a:hello b", "b"))
	// 将字符串的首尾字符代入匿名函数，如返回true，则该字符将被去除，否则保留
	fmt.Println(strings.TrimFunc("hello world", func(r rune) bool {
		if r == 'e' || r == 'd' {
			return true
		} else {
			return false
		}
	}))
}
