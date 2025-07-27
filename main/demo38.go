package main

import (
	"fmt"
	"sync"
	"time"
)

//单项管道

var wg sync.WaitGroup

func fn1(ch chan<- int) { //参数声明为只写管道
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v\n", i)
		time.Sleep(time.Millisecond * 50)
	}
	close(ch) //关闭管道
}

func fn2(ch <-chan int) { //参数声明为只读管道
	defer wg.Done()
	for v := range ch {
		fmt.Printf("读取数据%v\n", v)
		time.Sleep(time.Millisecond * 50)

	}
}

func main() {

	//只读只写的管道

	//只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 12
	//<-ch2 //invalid operation: cannot receive from send-only channel ch2 (variable of type chan<- int)

	//只读
	//ch3 := make(<-chan int, 2)

	//ch3 <- 23 //invalid operation: cannot send to receive-only channel ch3 (variable of type <-chan int)

	var ch = make(chan int, 10)

	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)

	wg.Wait()
	fmt.Println("主程序退出...")

}
