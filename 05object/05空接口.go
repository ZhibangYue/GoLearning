package main

func main051() {
	// data是一个切片，内容是空接口类型
	//interface{}是空接口
	data := make([]interface{}, 0)
	data = append(data, "字符串")
	data = append(data, 111)
	data = append(data, 1.23)

}
