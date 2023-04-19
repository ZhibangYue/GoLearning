package main

import (
	"bufio"
	"fmt"
	"os"
)

func main031() {
	// 使用追加写，不改变权限
	file, err := os.OpenFile(`D:\go_craft\GoLearning\06file\02ioutil文件读取.go`, os.O_WRONLY|os.O_APPEND, 0)
	// 创建写os.CREATE|os.O_WRONLY|os.O_APPEND
	// 覆盖写os.O_WRONLY|os.O_TRUC
	if err != nil {
		return
	}
	defer file.Close()
	// 创建文件的缓冲写出器
	writer := bufio.NewWriter(file)
	writer.WriteString("emmm")
	// 把缓冲区中的写入
	writer.Flush()
	fmt.Println("文件写出完毕")
}
