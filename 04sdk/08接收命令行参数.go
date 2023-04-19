package main

import (
	"fmt"
	"os"
)

func main081() {
	fmt.Println(os.Args)
	for i, arg := range os.Args {
		if arg == "rmb" {
			fmt.Println(os.Args[i+1])
		}
	}
}

// go build xxx 编译
// 08接收命令行参数 -r 124 rmb 1234
