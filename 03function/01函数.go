package main

import "fmt"

func SayHello(name string, n int) string {
	fmt.Println("hello，", name)
	for i := 0; i < n; i++ {
		fmt.Println("向你致以节日的问候")
	}
	return "谢谢"
}
func main011() {
	reply := SayHello("铁柱", 3)
	fmt.Println(reply)
}
