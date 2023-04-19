package main

// strings字符串处理包
import (
	"fmt"
	"strings"
)

/*
检索子串
Contains(s,substr string) bool
ContainsRune(s string, r int) bool
ContainsAny(s, chars string) bool
Index(s, sep string) int
Index(s, chars string) int
*/
func main011() {
	// 检索子串， 参数为两个字符串，返回布尔值
	fmt.Println(strings.Contains("abcde", "a")) // true
	// ContainsRune判断字符串中是否有参数r对应的字符，%d->%c
	// 123的字符形式是{
	fmt.Printf("%d的字符形式：%c\n", 123, 123)
	// {的UTF8码值（序号）形式是U+007B，就是十进制的123
	fmt.Printf("%s的字符形式：%U;\n", "{", '{')
	fmt.Println(strings.ContainsRune("abcde", 123)) // false
	// ContainsAny判断s中是否含有chars中的字符
	fmt.Println(strings.ContainsAny("hello", "abcde"))
}

func main012() {
	// 遍历字符串
	for i, c := range "你妹啊" {
		fmt.Printf("i=%d,c=%c\n", i, c)
	}
	// Index 子串sep在字符串s中第一次出现的位置， 不存在返回-1
	fmt.Println(strings.Index("hello", "el"))
	fmt.Println(strings.Index("hello", "h"))
	// 在UTF-8字符集中，一个汉字占三字节（1111,1111,1111,1111,1111,1111），此处的位置对应第几字节
	fmt.Println(strings.Index("你妹啊", "啊")) // 6
	// IndexAny 字符串chars中的任一UTF-8码值在s中第一次出现的位置，如果不存在或者chars为空字符串则返回-1
	fmt.Println(strings.IndexAny("你妹啊", "笑你妹")) // 0

}
