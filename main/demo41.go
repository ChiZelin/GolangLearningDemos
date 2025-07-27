package main

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
var count = 0

var wg sync.WaitGroup

var mutex sync.Mutex //声明互斥锁

func test() {
	mutex.Lock() //加锁
	count++
	time.Sleep(time.Millisecond * 100)
	fmt.Println("the count is: ", count)
	mutex.Unlock() //解锁
	wg.Done()
}

func main() {

	for r := 0; r < 50; r++ {
		wg.Add(1)
		go test()
	}

	wg.Wait()
	fmt.Println("主程序退出...")

}
