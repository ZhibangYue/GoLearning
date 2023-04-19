package main

import (
	"fmt"
	"os"
)

/*
Getwd() (dir string, err error) 返回当前工作目录的根路径
Getenv(key string) 返回环境变量的值，没有则返回空
Environ() []string 以键值对的形式返回所有环境变量
Hostname() (name string, err error) 返回当前设备在网络中的主机名
TempDir() 返回用户的临时文件夹
IsPathSeparator(c uint8) bool 判断是否为合法的路径分隔符，只能接受uint8以内的字符
Stat(name string) (FileInfo error)
*/

func main061() {
	//dir, err := os.Getwd()
	//fmt.Println(dir, err)
	//getenv := os.Getenv("Path")
	//fmt.Println(getenv)
	//fmt.Println(os.Environ())
	//name, err := os.Hostname()
	//fmt.Println(name, err)
	//fmt.Println(os.TempDir())
	//fmt.Println(os.IsPathSeparator('\\'))
	stat, err := os.Stat("D:/go_craft/GoLearning/04sdk/05strings-分割拼接.go")
	fmt.Println(stat, err)
	fmt.Println(stat.IsDir())
	fmt.Println(stat.Name())
	fmt.Println(stat.Size())
	fmt.Println(stat.Mode())
	fmt.Println(stat.ModTime())

}
