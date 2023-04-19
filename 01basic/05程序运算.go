package main

import (
	"fmt"
	"math"
)

// 数学运算
func main051() {
	// 四则运算
	fmt.Println("5+3=", 5+3)
	fmt.Println("5-3=", 5-3)
	fmt.Println("5*3=", 5*3)
	fmt.Println("5/3=", 5/3)
	fmt.Println("5%3=", 5%3)
	// 乘方和开方，整型只会得到整型的结果，开方时应设置为浮点数
	fmt.Println("5^3=", math.Pow(5, 3))
	fmt.Println("125^1/3", math.Pow(125.0, 1.0/3.0))
	// 四舍五入，负数的四舍五入是先对绝对值四舍五入再加负号
	fmt.Println("3.49->", math.Round(3.49))
	fmt.Println("-3.49->", math.Round(-3.51))
	// 纯舍和纯入
	fmt.Println("3.01->", math.Floor(3.01))
	fmt.Println("3.01->", math.Ceil(3.01))
	// 三角函数，参数必须是弧度而不是角度
	fmt.Println("sin30°=", math.Sin((30.0/180)*math.Pi))
	fmt.Println("cos30°=", math.Cos((30.0/180)*math.Pi))
	fmt.Println("tan30°=", math.Tan((30.0/180)*math.Pi))
	// 反三角函数
	fmt.Println("arcsin0.5=", math.Asin(0.5)*180/math.Pi)
}

// 逻辑运算，与或非
func main052() {
	a1 := (1+1 == 2)
	a0 := (1+1 != 2)
	fmt.Println(a1, a0, a1 && a0, a1 || a0, !a1)
}

/*
位运算
& 按位与， 都为1则为1
| 按位或， 有1则为1
^ 按位异或， 相同则为1
移位运算，<< 左移 >> 右移
*/
func main053() {
	a := 23 // 1,0111
	b := 34 // 10,0010
	// 按位与得到10，也就是2
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a << 2) // 101.1100
	fmt.Println(a >> 2) // 101
	var c int8 = 23
	var d uint8 = 23
	fmt.Println(c<<4, a<<4) // 左移的溢出，c是int8，a是int64，所以c左移4位时发生了溢出
	fmt.Println(c<<5, d<<5) // 1110,0000无符号为224，有符号为-32，首位符号位1表示负数补码，减1取反得32
}
