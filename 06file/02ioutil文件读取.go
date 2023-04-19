package main

import (
	"fmt"
	"os"
)

/*
ioutil包下的API进行便捷文件读取。Go 1.16后ioutil被弃用，功能迁移至io和os中。
*/
func main021() {
	// 读取指定文件，获得其原始字节，内部调用了文件的打开和关闭
	bytes, err := os.ReadFile(`D:\go_craft\GoLearning\06file\02ioutil文件读取.go`)
	if err != nil {
		return
	}
	str := string(bytes)
	fmt.Println(str)
}
