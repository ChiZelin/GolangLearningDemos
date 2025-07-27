package main

import (
	"fmt"
	"sort"
)

func main(){

	//切片
	var arr1 []int
	fmt.Printf("%v - %T  -  长度：%v\n", arr1, arr1, len(arr1))
	fmt.Println(arr1 == nil) //空的切片等于 nil

	var arr2 = []int{1, 2, 3}
	fmt.Printf("%v - %T  -  长度：%v\n", arr2, arr2, len(arr2))

	var arr3 = []int{1: 2, 2:4, 3: 6} //:前面的是索引
	fmt.Printf("%v - %T  -  长度：%v\n", arr3, arr3, len(arr3))

	//遍历和数组一样


	//基于数组定义切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[:] //获取数组里面的所有值

	fmt.Printf("%v - %T\n", b, b)

	c := a[1:4] //左闭右开
	fmt.Printf("%v - %T\n", c, c)

	d := a[:2] //获取下标是2之前的数据
	fmt.Printf("%v - %T\n", d, d)

	e := a[2:] //获取下标是2之后的数据
	fmt.Printf("%v - %T\n", e, e)


	//切片再切片 同理

	//切片的长度和容量
	var arr4 = []int {1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("长度%d  容量%d\n", len(arr4), cap(arr4)) //长度7  容量7

	f := arr4[2:]
	fmt.Printf("长度%d  容量%d\n", len(f), cap(f)) //长度5  容量5

	g := arr4[2:5]
	fmt.Printf("长度%d  容量%d\n", len(g), cap(g)) //长度3  容量5

	//make 函数创建一个切片  make([]T, size, cap)
	var sliceA = make([]int, 4, 8)

	fmt.Println(sliceA)
	fmt.Printf("长度%d  容量%d\n", len(sliceA), cap(sliceA)) //长度3  容量5

	sliceA[0] = 1
	fmt.Println(sliceA)

	//append 追加数据, 能自动扩容 当容量小于1024时 扩容时容量翻倍 大于等于1024时 按照1.25倍增长
	sliceA = append(sliceA, 2, 3, 4, 5, 6)
	fmt.Println(sliceA)
	fmt.Printf("长度%d  容量%d\n", len(sliceA), cap(sliceA)) //长度9  容量16

	var sliceB = []int{1, 2, 3}

	sliceA = append(sliceA, sliceB...) //固定语法 加... 表示合并切片
	fmt.Println(sliceA)
	fmt.Printf("长度%d  容量%d\n", len(sliceA), cap(sliceA)) //长度12  容量16



	//切片是引用数据类型， 可用copy()函数复制切片 copy的是值 是新开辟一块内存空间
	slice1 := []int{1, 2, 3, 4}
	slice2 := make([]int, 4, 4)
	copy(slice2, slice1)
	slice2[3] = 5
	fmt.Println(slice1) //[1 2 3 4]
	fmt.Println(slice2) //[1 2 3 5]

	fmt.Printf("%p\n", slice1)
	fmt.Printf("%p\n", slice2)

	//删除切片中的数据 没有专门方法 用 append 方法实现
	// 删除切片中索引为2的元素
	slice1 = append(slice1[:2], slice1[3:]...)
	fmt.Println(slice1)

	//字符串不能修改转换成 可转换成切片修改 然后再转回字符串
	s2 := "您好 golang"
	runeStr := []rune(s2)
	fmt.Println(runeStr)
	runeStr[0] = '大'
	s2 = string(runeStr)
	fmt.Println(s2)


	//排序 可用 sort 包中的方法
	intlist := []int{2, 4, 3, 5, 7, 6}
	fmt.Println(intlist)
	sort.Ints(intlist)
	fmt.Println(intlist)
	//降序
	sort.Sort(sort.Reverse(sort.IntSlice(intlist))) //降序
	fmt.Println(intlist)

}