package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fn1(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch) //关闭管道
}

func fn2(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("读取数据%v\n", v)
		time.Sleep(time.Millisecond * 50)

	}
}

func main() {

	var ch = make(chan int, 10)

	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)

	wg.Wait()
	fmt.Println("主程序退出...")

}
