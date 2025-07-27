package main

import (
	"fmt"
)

func main() {

	//管道 channel  是引用数据类型  改变副本会影响原来变量的值

	ch := make(chan int, 3)

	ch <- 10
	ch <- 21
	ch <- 32

	a := <-ch

	fmt.Println(a) //10

	<-ch //从管道中取值21 但没赋值
	c := <-ch
	fmt.Println(c) //32

	ch <- 56
	fmt.Printf("值：%v 容量：%v 长度%v\n", ch, cap(ch), len(ch)) //值：0xc0000ae080 容量：3 长度1

	// 管道阻塞 在没有协程的情况下 满了之后还往里面加数据则阻塞
	ch <- 10
	ch <- 21
	//ch <- 32 //fatal error: all goroutines are asleep - deadlock!

	ch2 := make(chan string, 1)

	ch2 <- "aaa"
	<-ch2
	//<-ch2 //fatal error: all goroutines are asleep - deadlock!   在没有协程的情况下 管道中没有值了 再取也阻塞

	//遍历管道

	var ch3 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch3 <- i
	}

	// for i := 1; i <= 10; i++ { // 通过for 循环遍历时 不用关闭
	// 	fmt.Println(<-ch3)
	// }
	close(ch3)

	for v := range ch3 { // for range 遍历前 管道必须关闭
		fmt.Println(v)
	}

}
