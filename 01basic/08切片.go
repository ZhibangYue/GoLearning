package main

import "fmt"

// 切片的长度与容量
// cap(slice)获得切片的容量
// 创建之初，容量等于长度
// 扩张时，一旦容量无法满足需要，就以翻倍的策略扩容
func main081() {
	var slice = []int{1, 2, 3}
	fmt.Println("len(slice)=", len(slice), "cap(slice)=", cap(slice))
	slice = append(slice, 4)
	fmt.Println("len(slice)=", len(slice), "cap(slice)=", cap(slice))
	slice = append(slice, 5)
	fmt.Println("len(slice)=", len(slice), "cap(slice)=", cap(slice))
}

// 切片的兼并
func main082() {
	var slice1 = []int{1, 2, 3}
	slice2 := []int{4, 5, 6, 7}
	// 变量不能声明而不使用
	// 下划线起这种作用
	for _, value := range slice2 {
		slice1 = append(slice1, value)
	}
	fmt.Println(slice1)
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}

// 切片的创建
func main083() {
	// 创建长度为3的切片
	slice1 := make([]int, 3)
	// 创建长度为3容量为5的切片
	slice2 := make([]int, 3, 5)
	fmt.Println(slice1, slice2, cap(slice1), cap(slice2))
}

// 切片不是拷贝，而是地址的引用
func main084() {
	var array = [5]int{0, 1, 2, 3, 4}
	slice1 := array[:]
	slice2 := slice1[:]
	fmt.Println(array, slice1, slice2)
	fmt.Println(&array[0], &slice1[0], &slice2[0])
	// 对切片的修改会影响所有的
	slice1[1] = 11
	fmt.Println(array, slice1, slice2)
	// 切片扩容时，地址的引用被断开了
	slice2 = append(slice2, 5, 6, 7)
	fmt.Println(slice1, array, slice2)
	fmt.Println(&array[0], &slice1[0], &slice2[0])

}
