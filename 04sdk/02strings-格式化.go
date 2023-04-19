package main

import (
	"fmt"
	"strings"
)

/*
格式化
ToUpper(s) 转换为全大写
ToLower(s) 转换为全小写
ToTitle(s) 转换为标题格式
*/

func main021() {
	fmt.Println(strings.ToUpper("hELlo"))
	fmt.Println(strings.ToLower("hELlo"))
	fmt.Println(strings.ToTitle("hELlo"))
}
