package main

import (
	"fmt"
)

// Sub01 无返回值
func Sub01(a, b int) {
	ret := a - b
	fmt.Println("a-b=", ret)
}

// Sub02 有返回值，返回a-b的结果(int)
func Sub02(a, b int) int {
	ret := a - b
	return ret
}

// Sub03 有返回值，返回a-b的结果(int)，起名为difference
func Sub03(a, b int) (difference int) {
	difference = a - b
	return difference
}

// Sub04 多个返回值
func Sub04(a, b int) (sum int, difference int, product int, quotient float64) {
	sum = a + b
	difference = a - b
	product = a * b
	quotient = float64(a) / float64(b)
	return

}
func main031() {
	Sub01(5, 3)
	ret := Sub03(5, 3)
	fmt.Println(ret)
	sum, difference, product, quotient := Sub04(3, 5)
	fmt.Println(sum, difference, product, quotient)
}
