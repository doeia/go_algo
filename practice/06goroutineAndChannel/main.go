package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	for i := 0; i < 10; i++ {
		// fmt.Println("2")
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	//ch := make(chan<- int, 3) writeonly
	//ch := make(<-chan int, 3) readonly
	ch := make(chan int, 3) // 管道是引用类型，先进先出，超过长度将管道阻塞(deadlock)
	ch <- 10
	a := <-ch
	fmt.Println(a)

	wg.Add(1)
	go test()

	// 问题：主进程完成就结束了。不会等协程完毕,用WaitGroup解决
	for i := 0; i < 10; i++ {
		// fmt.Println("1")
		time.Sleep(time.Millisecond * 20)
	}
	wg.Wait()
	fmt.Println("finish main")
}
