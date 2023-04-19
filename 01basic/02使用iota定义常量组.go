package main

import "fmt"

/*
const(

	USA = 0
	China = 1
	Russia = 2
	Britain = 3
	France = 4

)
*/

/*
使用iota定义一组常量，iota默认是0，Untyped int.
后边的常量会自动沿用第一个的表达式（但iota的值会逐个递增）
*/
const (
	USA = iota + 1
	China
	Russia
	Britain
	France
)

func main04() {
	fmt.Println(USA, China, Russia, Britain, France)
}
