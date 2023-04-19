package main

import "fmt"

func main061() {
	a := 12
	aPtr := &a
	fmt.Printf("a的地址是%p\n", &a)
	fmt.Printf("a的地址是%p\n", aPtr)
	fmt.Println("aPtr是", aPtr)
	fmt.Println("aPtr的值是", *aPtr)
	var xPtr *int
	fmt.Println("空指针xPtr是", xPtr) // nil
	x := 789
	xPtr = &x
	fmt.Println(x, xPtr, *xPtr)
	*xPtr = 456
	fmt.Println(x)
}

// 指针是严格检测数据类型的
func main062() {
	//var xPtr *int
	//var y = "string"
	//xPtr = &y
}

// 指针的指针
func main063() {
	var x = 123
	var xPtr = &x
	var xPP = &xPtr
	fmt.Printf("xPtr的类型是%T,xPP的类型是%T\n", xPtr, xPP)
	fmt.Println(xPtr, *xPP, xPP, x, *xPtr, **xPP)
}
