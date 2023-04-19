package main

import (
	"fmt"
	"os"
)

func main051() {
	fileinfo, err := os.Stat("D:\\go_craft\\GoLearning\\06file\\04ioutil文件写出.txt")
	if err == nil && fileinfo != nil {
		fmt.Println("文件存在")
	} else if os.IsNotExist(err) {
		fmt.Println("文件不存在")
	} else {
		fmt.Println("文件异常", fileinfo, err)
	}
}
