package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main051() {
	// 当前时间距离1970年0时所逝去的时间纳秒数
	seed := time.Now().UnixNano()
	n := rand.Intn(1000)
	fmt.Println(n)
	myRandom := rand.New(rand.NewSource(seed))
	fmt.Println(myRandom.Intn(1000))
}
