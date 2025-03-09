package mutex

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup

var mutext sync.Mutex

func test() {
	mutext.Lock()
	count++
	fmt.Println("this count is ", count)
	time.Sleep(time.Millisecond)
	mutext.Unlock()
	wg.Done()
}

func main() {
	// 开启多个协程会抢占资源，需要加锁
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}
