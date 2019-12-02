package main

import (
	"fmt"
	"time"
)

func test(ch chan int) {

	go func() {
		ch <- 100
		fmt.Printf("入队后的数据 %v\n", ch)
		ch <- 200
		fmt.Printf("入队后的数据 %v\n", ch)
		ch <- 300
		fmt.Printf("入队后的数据 %v\n", ch)
	}()
}

// 生产者消费者模型
func producer(ch chan<- int) { // 只写
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Send:", i)
	}
}

func consumer(ch <-chan int) {
	for i := 0; i < 5; i++ {
		v := <-ch
		fmt.Println("Receive:", v)
	}
}
func main() {
	//var ch chan int
	//ch = make(chan int)
	//test(ch)
	//data1 := <-ch
	//data2 := <-ch
	//data3 := <-ch
	//fmt.Println(data1, data2, data3)
	//
	//var ch1 chan<- int // 只写 channel
	//var ch2 <-chan int // 只读
	ch := make(chan int) // 非缓冲模式
	go producer(ch)
	go consumer(ch)

	time.Sleep(3 * time.Second)
}
