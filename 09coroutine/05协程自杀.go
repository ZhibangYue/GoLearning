package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
runtime.Goexit()使得协程自杀
而当主协程自杀后，子协程就自由了
*/
func main051() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i == 5 {
				fmt.Println("goroutine died!")
				runtime.Goexit()
			}
			fmt.Println("goroutine", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()
	for i := 1; i <= 10; i++ {
		fmt.Println("main", i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("main over")
}
