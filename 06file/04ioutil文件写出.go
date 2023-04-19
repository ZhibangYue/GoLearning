package main

import (
	"fmt"
	"os"
)

func main041() {
	// 注意写出的内容应为原始字节
	err := os.WriteFile("D:\\go_craft\\GoLearning\\06file\\04ioutil文件写出.txt", []byte("写入内容"), 0)
	if err != nil {
		return
	}
	fmt.Println("写出成功")
}
