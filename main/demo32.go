package main

//goroutine

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup //用 sync.WaitGroup 来确保协程执行完毕后主进程才退出

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好Golang", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	go test() //表示开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("你好Golang", i)
		time.Sleep(time.Millisecond * 20)
	}
	wg.Wait()
	fmt.Println("主线程退出")
}
