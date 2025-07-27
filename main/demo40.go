package main

import (
	"fmt"
	"sync"
	"time"
)

//goroutine 中 panic 处理

var wg sync.WaitGroup

func sayHello() {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Printf("hello, world\n")
	}

	wg.Done()
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
		wg.Done()
	}()

	var myMap map[int]string
	myMap[0] = "golang" // error
}

func main() {

	wg.Add(1)
	go sayHello()
	wg.Add(1)
	go test()

	wg.Wait()
	fmt.Println("主程序退出...")

}
