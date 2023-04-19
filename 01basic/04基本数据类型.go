package main

import "fmt"

// 补码：用来表示负数。原码取反再加一，3的原码0000,0011；补码1111,1101

// 基本数据类型
func main041() {
	// 整型 int8、int16、int32、int64、unit
	var a byte = 255
	var b rune = 1
	// int类型不指定位数时跟机器自身有关，如果是32位那么默认为int32，64位则是int64,uint同理
	var c int = -1
	var d uint = 1
	fmt.Println(a, b, c, d)
	// 浮点型包括float32和float64。
	// 复数包括complex64和complex128，实部和虚部分别是32位/64位的。
	// 布尔类型 true false
	var isClever bool = true
	fmt.Printf("isClever的类型是%T\n", isClever)
	var hp = 1.23
	fmt.Printf("hp的类型是%T\n", hp)
}

/*
基本类型占位符
%s 字符串占位符
%d 整型占位符（十进制）
%f 浮点型占位符， %.2f精确到小数点后两位
%t 布尔占位符
右对齐，左对齐，%5s，%-5s等
*/
func main042() {
	fmt.Printf("%6s是名字，%.2f是数字，这个是在%-5s空格", "name", 1.23333, "后面")
}
