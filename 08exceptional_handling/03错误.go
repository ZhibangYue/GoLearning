package main

import (
	"errors"
	"fmt"
	"math"
)

/*
相较于恐慌，错误是一种相对温和的异常解决方案，函数会把结果和错误同时返回
如果结果正确时错误为空，如果错误不为空时结果为空（或没有业务意义的默认值）
*/

func GetCircleArea(radius float64) (float64, error) {
	if radius < 0 {
		err := errors.New("半径不能为负数")
		return 0, err
	}
	return math.Pi * radius * radius, nil
}

func main031() {
	area, err := GetCircleArea(-5)
	if err != nil {
		fmt.Println("面积计算失败", err)
		return
	}
	fmt.Println(area)
	fmt.Println("其他功能")
}
