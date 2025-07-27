package main

import "fmt"

//泛型

func main() {
	// 原文件路径
	strs := []string{"aaa", "bbb"}
	is := []int{1, 2}
	fs := []float64{1.2, 2.1}

	printArray(strs)
	printArray(is)
	printArray(fs)

	printArray2(strs)
	printArray2(is)
	printArray2(fs)

}

func printArray[T string | int | float64](arr []T) { //指定约束任何类型
	for _, a := range arr {
		fmt.Println(a)
	}
}

func printArray2[T any](arr []T) { //约束任何类型 any
	for _, a := range arr {
		fmt.Println(a)
	}
}
