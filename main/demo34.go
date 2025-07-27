package main

//开启 5 个协程 分别打印 5 条数据

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("协程(%v)打印的第%v条数据\n", num, i+1)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go test(i + 1)
	}

	wg.Wait()
	fmt.Printf("主线程退出\n")
}
