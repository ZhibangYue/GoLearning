package main

import (
	"fmt"
	"time"
)

/*
CSP并发理论
异步async并行：多个任务并发执行
同步sync串行：多个任务依次执行
阻塞block：某个并发任务由于拿不到资源无法干活，从而无所事事地干等

进程并发-线程并发-协程并发
进程并发：CPU多核并发，不需切换时间片
线程并发：物理意义上的并发-CPU切换线程（伪并发）
协程并发：逻辑意义上的并发-一饰多角

异步回调async callback
A线程唤起B线程，令其干活
同时给B一个回调函数
命令B在干完活以后，执行这个回调函数
这个回调函数会与A线程发生交互
A不必阻塞等待B执行的结果，AB两个线程可以并发执行
利弊:1.效率高.2.回调地狱CallbackHell,逻辑线不清晰

共享内存
多个并发线程通过共享内存的方式交互数据
线程安全问题:AB间共享的数据地址可能被C并发修改

同步锁/资源锁
为了解决共享内存所导致的线程安全问题，共享的内存地址在特定时间段被特定线
程锁定
加锁期间，其它线程无法访问，带来低效率问题

死锁
A锁住B的资源，B锁住A的资源，AB同时阻塞

线程池
背景：线程开销巨大
内存：保存上下文数据
CPU：调度
消耗CPU、内存
为了避免无度创建线程（内存溢出Out Of Memory），在一个池中创建一堆线程，循环使用这些线程，用完了以后重置丢回池中。
利：避免了无度创建线程，降低了OOM的风险；弊：无论如何都占用了一大块内存，而且线程数有上限

线程并发的弊端
开线程占内存，线程切换占CPU，共享内存不安全，加锁效率低下，回调地狱开发难度高
开销大、效率低、难度高

CSP模型
Communicating Sequential Process
可通信的序列化进程
并发的进程间通过管道进行通信

管道
最早由CSP模型提出
以点对点管道代替内存共享实现并发进程间的数据交互
相比内存共享数据交互的效率高很多
管道是一块内存，可以直接做逻辑上的并发调度，省去了CPU+内存的物理开销

协程
cooperate 合作
routine 事务
coroutine 协程（可以合作的并发任务）
gocoroutine
*/

func SayHello() {
	fmt.Println("通知主线程活干完了")
}

func main011() {
	go func(f func()) {
		time.Sleep(time.Second)
		for i := 11; i < 20; i++ {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
		f()
	}(SayHello)
	for i := 1; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}
