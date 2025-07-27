package main

import (
	"fmt"
	"sync"
	"time"
)

//用协程和管道 实现 判断 1-120000 中的素数

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}

	close(intChan)
	wg.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	for num := range intChan {
		var flag = true

		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}

	}

	wg.Done()
	exitChan <- true
}

func printPrime(primeChan chan int) {

	for v := range primeChan {
		fmt.Println(v)
	}

	wg.Done()
}

func main() {

	start := time.Now().Unix()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 1000)
	exitChan := make(chan bool, 16)

	wg.Add(1)
	go putNum(intChan)

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	wg.Add(1)
	go printPrime(primeChan)

	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan
		}

		close(exitChan)
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()

	end := time.Now().Unix()
	fmt.Println("执行完毕，退出程序...", end-start, "秒")

}
