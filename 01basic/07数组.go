package main

import "fmt"

// 数组的声明与赋值
func main071() {
	var a1 [10]int
	fmt.Println(a1)
	// [0 0 0 0 0 0 0 0 0 0] 自动补零
	var a11 = [5]int{1, 2, 3}
	fmt.Println(a11)

	var a2 [3]int = [3]int{1, 2, 3}
	fmt.Println(a2)
	var a3 = [3]int{}
	fmt.Println(a3)
	var a4 = [...]int{1, 2, 3}
	fmt.Printf("a4的类型是%T,值是%v", a4, a4)
}

// 数组的索引与长度
func main072() {
	a1 := [...]int{1, 2, 3, 4, 5}
	fmt.Println("第一个元素", a1[1])
	fmt.Println("长度是", len(a1))
}

// 数组的遍历
func main073() {
	a1 := [...]int{1, 2, 3, 4, 5}
	for index, value := range a1 {
		fmt.Printf("a[%d]=%v\n", index, value)
	}
	for i := 0; i < len(a1); i++ {
		fmt.Printf("a1[%d]=%d\n", i, a1[i])
	}
}

// 多维数组
func main074() {
	var arr [3][4]int
	arr[0] = [4]int{1, 2, 3}
	arr[1] = [4]int{1, 2, 3, 4}
	arr[2] = [4]int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(arr[0][3])
}
