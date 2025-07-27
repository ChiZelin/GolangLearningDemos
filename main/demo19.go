package main

import (
	"fmt"
	"time"
)


func main(){
	//Time

	timeObj := time.Now()
	fmt.Println(timeObj)
	
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("%d-%02d-%d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

	str := timeObj.Format("2006-01-02 15:04:05") //用 golang 诞生时间当模板
	fmt.Println(str)


	unixtime := timeObj.Unix()//获取当前时间戳
	fmt.Println("当前时间戳：", unixtime)


	//时间戳转换成时间
	var timesdpt int64 = 1752482204
	timeobj2 := time.Unix(timesdpt, 0) //0 代表纳秒部分
	var str2 = timeobj2.Format("2006-01-02 15:04:05")
	fmt.Println(str2)

	//日期字符串转换成时间戳
	var str3 = "2025-07-14 16:36:44"
	var tmp = "2006-01-02 15:04:05"
	timeObj3, _ := time.ParseInLocation(tmp, str3, time.Local)
	fmt.Println(timeObj3) //2025-07-14 16:36:44 +0800 CST
	fmt.Println(timeObj3.Unix()) //1752482204


	//time 包中定义的时间间隔类型的常量
	fmt.Println(time.Millisecond)
	fmt.Println(time.Second)
	fmt.Println(time.Hour)

	//时间操作函数
	var timeObj4 = time.Now()
	fmt.Println(timeObj4) //2025-07-14 16:58:30.6480737 +0800 CST m=+0.017524901
	timeObj4 = timeObj4.Add(time.Hour)
	fmt.Println(timeObj4) //2025-07-14 17:58:30.6480737 +0800 CST m=+3600.017524901


	//定时器
	ticker := time.NewTicker(time.Second) //定义一个1秒间隔的定时器
	n := 0
	for i := range ticker.C {
		fmt.Println(i)
		n++
		if n > 5 {
			ticker.Stop()
			break
		}
	}

	fmt.Println("aaa1")
	time.Sleep(time.Second * 5) // Sleep
	fmt.Println("aaa2")


}