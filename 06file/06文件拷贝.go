package main

import (
	"fmt"
	"io"
	"os"
)

func main061() {
	bytes, _ := os.ReadFile("D:\\go_craft\\GoLearning\\06file\\04ioutil文件写出.txt")
	err := os.WriteFile("D:\\go_craft\\GoLearning\\06file\\04ioutil文件写出2.txt", bytes, 0666)
	if err != nil {
		return
	}
	fmt.Println("拷贝成功")
}

/*
io.Copy() 参数为目标文件和源文件，返回字节数和错误
*/
func main062() {
	srcFile, _ := os.Open("D:\\go_craft\\GoLearning\\06file\\04ioutil文件写出.txt")
	dstFile, err := os.OpenFile("D:\\go_craft\\GoLearning\\06file\\04ioutil文件写出3.txt", os.O_CREATE|os.O_WRONLY, 0666)
	defer func() {
		srcFile.Close()
		dstFile.Close()
	}()
	write, err := io.Copy(dstFile, srcFile)
	if err != nil {
		return
	}
	fmt.Println("写出成功", write)
}
