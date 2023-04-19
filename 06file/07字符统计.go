package main

import (
	"fmt"
	"os"
)

func main071() {
	bytes, _ := os.ReadFile("D:\\go_craft\\GoLearning\\06file\\04ioutil文件写出.txt")
	str := string(bytes)
	for _, c := range str {
		fmt.Printf("%T %c %U\n", c, c, c)
	}
}
