package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
管道channel提供了一种通信机制。通过它，一个协程可以给另一个协程发送消息。
管道本身还关联了一个数据类型，即管道能发送的数据类型。
管道其实就是一片连续的内存，赋值或参数传递时传递的是指针。
*/

// 管道的创建和读写
func main061() {
	var myChan chan int
	fmt.Println(myChan)
	myChan = make(chan int)
	fmt.Println(myChan)
	// 向管道写入数据，由于此管道缓存能力为0，无法写入，因而形成死锁，导致阻塞。(管道装满，无法写入，死锁)
	//myChan <- 13
	myChan = make(chan int, 1)
	myChan <- 123
	// 从管道读取数据
	x := <-myChan
	fmt.Println(x)
	fmt.Println("main over")
}

// 管道的读写与死锁
func main062() {
	myChan := make(chan string, 3)
	// 缓存满了，没有人读，写入阻塞。已经读空，没有人写，读取阻塞。
	go func() {
		for i := 0; i < 5; i++ {
			x := <-myChan
			fmt.Println(x, "已读取")
		}
	}()
	for i := 0; i < 5; i++ {
		myChan <- strconv.Itoa(i)
		fmt.Println(i, "已写入")
		time.Sleep(1 * time.Second)
	}
}

/*
管道的关闭
关闭的管道可以读，不能写
未初始化的管道不能被关闭（因为其实是nil），否则会产生恐慌
不能重复关闭管道，否则也会产生恐慌
*/
func main063() {
	bools := make(chan bool, 3)
	go func() {
		time.Sleep(3 * time.Second)
		x := <-bools
		fmt.Println(x)
		x = <-bools
		fmt.Println(x)
		x = <-bools
		fmt.Println(x) // false
		// 当管道已为空，再读取会读取出零值
	}()
	bools <- true
	bools <- true
	close(bools)
	// 关闭后的管道无法写入，强行写入会产生恐慌
	// bools <- true

	for {
		time.Sleep(1 * time.Second)
	}
}

// 管道的遍历
func main064() {
	myChan := make(chan int, 5)
	go func() {
		for x := range myChan {
			fmt.Println("开始读取")
			fmt.Println(x)
		}
		// 即使被读取完，遍历仍在继续
		// 如果管道未关闭，会永远尝试下一次，而一旦管道关闭，阻塞遍历会自动结束
		fmt.Println("routine over")
	}()
	myChan <- 123
	myChan <- 124
	myChan <- 125
	// 管道关闭时会向所有读管道的协程发送通知，协程就会结束阻塞读取尝试
	close(myChan)
	for {
		time.Sleep(3 * time.Second)
	}

}
