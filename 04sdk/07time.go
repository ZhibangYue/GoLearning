package main

import (
	"fmt"
	"time"
)

/*
Now() 获取当前时间对象
Date() 创建时间对象
Sub() 计算时间差，负数表示将来
ParseDuration() 解析时间差
*/

func main071() {
	NowTime := time.Now() // 获取当前时间对象
	fmt.Println(NowTime.Year())
	fmt.Println(NowTime.YearDay())
	fmt.Println(NowTime.Month())
	fmt.Println(NowTime.Day())
	fmt.Println(NowTime.Date())
	fmt.Println(NowTime.Weekday())
	fmt.Println(NowTime.Hour())
	fmt.Println(NowTime.Minute())

	dateTime := time.Date(2023, time.January, 25, 12, 0, 0, 0, time.Now().Location())
	fmt.Println(dateTime)

	fmt.Println(NowTime.Sub(dateTime))
	duration, _ := time.ParseDuration(string(NowTime.Sub(dateTime)))
	newTime := NowTime.Add(duration)
	fmt.Println(newTime, duration)
}
