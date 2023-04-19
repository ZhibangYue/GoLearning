package main

import (
	"fmt"
	"math"
)

/*
type error interface{
	Error() string
}
*/

/*
实现上面的接口就可以实现自定义异常
*/

type InvalidRadiusError struct {
}

func (ire *InvalidRadiusError) Error() string {
	return "半径非法"
}

func GetCircleArea2(radius float64) (float64, error) {
	if radius < 0 {
		//err := errors.New("半径不能为负数")
		//return 0, new(InvalidRadiusError)
		return 0, &InvalidRadiusError{}
		// 因为是对象的指针实现了Error方法，所以这里取指针
	}
	return math.Pi * radius * radius, nil
}

func main041() {
	area2, err := GetCircleArea2(-5)
	fmt.Println(area2, err)
}
