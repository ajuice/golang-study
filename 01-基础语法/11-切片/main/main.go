package main

import "fmt"

func sliceTest() {
	// var s [5]int 数组

	// 切片是引用类型，可扩容
	//var s1 []int // 切片，切片与数组的区别为 [] 里面是否由数字
	//s1[1] = 10
	// make 声明
	s2 := make([]int, 5, 10) // 声明一个个数为5 容量为10 的切片

	fmt.Printf("%d,%d, %v\n", s2, len(s2), cap(s2)) // cap 获取容量
	// len 略
	// cap 略
	// append 同string
	// copy
	fmt.Println("--------------")
	s3 := []int{1, 2, 3}
	//s4 := []int{4, 5, 6}
	s5 := make([]int, 10)
	fmt.Println(copy(s5, s3), s3, s5)
}

func main() {
	sliceTest()
}
