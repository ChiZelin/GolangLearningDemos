package main

//goroutine

import (
	"fmt"
	"runtime"
)

func main() {

	cpuNum := runtime.NumCPU() //获取当前计算机上面的cup个数

	fmt.Println("cpuNum=", cpuNum)

	runtime.GOMAXPROCS(cpuNum - 1) //设置使用多少个cpu

	fmt.Println("OK")

}
