package goroutineandchannel

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("writing", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch)
	wg.Done()
}
func fn2(ch chan int) {
	for v := range ch {
		fmt.Println(v)
		fmt.Println("reading", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}

func main() {
	var ch = make(chan int, 10)

	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)

	wg.Wait()
	fmt.Println("finish")

	// 多路复用　select
	// for{
	// 	select{
	// 	case  v := <-intChan:
	// 		fmt.Println("int")
	//		time.Sleep(time.Millisecond * 50)
	// 	case  v := <-stringChan:
	// 		fmt.Println("string")
	//		time.Sleep(time.Millisecond * 50)
	// 	default:
	// 		fmt.Println("finish")
	//		return
	// 	}
	// }
}
