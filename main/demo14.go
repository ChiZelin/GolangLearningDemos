package main

import "fmt"

var a = "全局变量"

func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

func sumFn2(x ...int) int { //无限参数写法
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}

func subFn(x, y int) int {  //参数类型如果一样可以省略 只在后面保留一个
	sub := x - y
	return sub
}

//return 关键字一次可以返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//return 关键字一次可以返回多个值， 返回值变量名可直接声明
func calc2(x, y int) (sum int, sub int) {
	sum = x + y //此时就直接赋值 而不用类型推导了
	sub = x - y
	return
}

//自定义一个方法类型
type calcType func(int, int) int

func calc3(x, y int, cb calcType) int { //可传入函数类型 sumFn subFn
	return cb(x, y)
}


//函数也可以作为另一个函数的返回值
func do(o string) calcType {
	switch o {
	case "+":
		return sumFn
	case "-":
		return subFn
	default:
		return nil
	}

}

func main() {

	//函数
	sum1 := sumFn(1, 2)
	fmt.Println(sum1)

	sub := subFn(5, 3)
	fmt.Println(sub)

	sum2 := sumFn2(1, 3, 4)
	fmt.Println(sum2)

	sum3, sub3 := calc2(1, 5) //调用函数 一次返回多个值
	fmt.Printf("sum:%v sub:%v\n", sum3, sub3)


	fmt.Println(a)


	sub4 := calc3(10, 15, subFn)
	sum4 := calc3(10, 15, sumFn)

	fmt.Printf("sum:%v sub:%v\n", sum4, sub4)

	//传入匿名方法
	m := calc3(10, 15, func(x, y int) int {
		return x * y
	})

	fmt.Println(m)

	d := do("+")
	r := d(10, 5)

	fmt.Println(r)


	//匿名自执行函数, 注意后面有括号才能执行
	func(){
		fmt.Println("执行了")
	}()

	//函数赋值给变量
	var fn = func(){
		fmt.Println("执行了2")
	}

	fn()
}
