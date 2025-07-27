package main

import (
	"fmt"
	"time"
)

//select 多路复用

func main() {

	intChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//close(intChan)  //使用 select 时管道不能关闭
	stringChan := make(chan string, 10)

	for i := 0; i < 10; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//close(stringChan) //使用 select 时管道不能关闭
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 读取数据 %d\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取数据 %v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Printf("数据获取完毕...")
			return
		}

	}

}
