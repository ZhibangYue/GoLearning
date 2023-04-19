package main

import "fmt"

/*
字符串比大小
比首字符在字符集中出现的序号，首字符相同则比较第二个，依此类推
go的编译器默认使用UTF8字符集
0123456789
*/
func main011() {
	var birthday string
	fmt.Println("请输入您的生日，例如0823")
	fmt.Scan(&birthday)
	// if-elseif-else选择结构
	if birthday >= "0823" && birthday <= "0922" {
		fmt.Println("是处女座")
	} else if birthday >= "0923" && birthday <= "1022" {
		fmt.Println("是天秤座")
	} else {
		fmt.Println("不是处女座，不是天秤座")
	}
}

/*
选择结构switch
case单点或单点集
*/
func main012() {
	fmt.Println("请输入出生月份（1-12）\n")
	var month int
	fmt.Scan(&month)
	fmt.Printf("month=%d", month)
	switch month {
	case 1:
		fmt.Println("你是魔羯座")
	case 2:
		fmt.Println("你是水瓶座")
	case 3, 4, 5:
		fmt.Println("你出生在春天")
	// default 表示没有落在上述任一情况
	default:
		fmt.Println("不是魔羯座不是水瓶座")
	}
}

/*
选择结构switch
case区间
*/
func main013() {
	fmt.Println("请输入出生月份（1-12）\n")
	var month int
	fmt.Scan(&month)
	fmt.Printf("month=%d", month)
	switch {
	case month >= 6 && month <= 8:
		fmt.Println("是夏天")
	case month >= 9 && month <= 11:
		fmt.Println("是秋天")
	case month >= 3 && month <= 5:
		fmt.Println("是春天")
	case month == 12 || month == 1 || month == 2:
		fmt.Println("是冬天")
	// default 表示没有落在上述任一情况
	default:
		fmt.Println("你出生在火星")
	}
}
