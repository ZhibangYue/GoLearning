package main

import (
	"flag"
	"fmt"
)

/*
go build -o xxx.exe xxx.go 编译
go run xxx.go 编译并执行，生成的exe在临时目录下
*/

func main091() {
	// 定义一个字符串类型的命令行参数：参数名称name，参数默认值无名氏，参数用途说明“你妹芳名”。
	// 得到一个将来存储name参数的值的地址/指针
	// 传入未定义或格式错误的参数时会提示
	// 布尔型传参时需要用=，即-bool=false
	strPtr := flag.String("name", "无名氏", "你妹芳名")
	agePtr := flag.Int("age", 0, "你妹芳龄")
	// 两种方法是等效的
	// 但StringVar（或IntVar等其他的）要传入地址&string，不能传入空指针，因为空指针nil对应着地址0x0，内容是不可修改的
	var aaa string
	flag.StringVar(&aaa, "aaa", "is a", "这是a")
	// 解析flag获取参数值，必须被调用。
	flag.Parse()
	fmt.Println(*strPtr, *agePtr, aaa)
}
