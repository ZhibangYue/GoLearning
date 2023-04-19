package main

import (
	"fmt"
	"strings"
)

/*
字符串比大小
比较首字母顺序，大写在前小写在后
Compare(a,b string)
a>b 1
a<b -1
a=b 0
*/
func main031() {
	fmt.Println(strings.Compare("hello", "asdasd"))
	fmt.Println(strings.Compare("hello", "jsdasd"))
	fmt.Println(strings.Compare("hello", "hello"))
	fmt.Println(strings.Compare("hello", "HELLO"))
}
