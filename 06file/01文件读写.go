package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
buffer 缓冲区
util(utility) 便利的工具
perm(permission) 权限
*/

// 文件的打开与关闭
func main011() {
	// 打开文件，返回文件指针和错误
	filePtr, err := os.Open("D:\\go_craft\\GoLearning\\06file\\01文件读写.go")
	// 如果有错误，说明打开失败，return结束函数
	if err != nil {
		fmt.Println("文件打开失败,err=", err)
		return
	} else {
		fmt.Println("文件打开成功")
	}
	// 文件是IO资源，使用完毕一定要关闭，以释放IO资源
	defer func() {
		filePtr.Close()
		fmt.Println("文件成功关闭")
	}()
	fmt.Println(filePtr)
}

/*
OpenFile(name string, flag int, perm FileMode) (*File, error)
返回文件指针和错误
Open函数是OpenFile的特例
flag见下
perm表示permission，权限4表示读取，2表示写入。1表示可执行。rwx。0666，0755
*/
//	O_RDONLY // open the file read-only.
//	O_WRONLY  // open the file write-only.
//	O_RDWR    // open the file read-write.
//	// The remaining values may be or'ed in to control behavior.
//	O_APPEND // append data to the file when writing.
//	O_CREATE // create a new file if none exists.
//	O_EXCL  // used with O_CREATE, file must not exist.
//	O_SYNC  // open for synchronous I/O.
//	O_TRUNC  // truncate regular writable file when opened.

/*
缓冲区读取
*/
func main012() {
	file, _ := os.Open("D:/go_craft/GoLearning/06file/01文件读写.go")
	// 创建文件的缓冲读取器
	reader := bufio.NewReader(file)
	defer func() {
		file.Close()
		fmt.Println("文件成功关闭")
	}()
	// 遍历读取每一行
	for {
		// dilim表示定界符，它的意思是读取到\n，即换行符，也就是读取一行
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("文件读取失败,err=", err)
			if err == io.EOF {
				fmt.Println("已到文件末尾")
				break
			}
		} else {
			fmt.Print(line)
		}
	}
}
