package main

import (
	"fmt"
	"time"
)

func main031() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("努力奋斗了%d个小时", i)
		time.Sleep(500 * time.Millisecond)
		if i >= 20 {
			//break
			goto GAMEOVER
		}
	}

GAMEOVER:
	fmt.Println("GAME OVER!")
}
