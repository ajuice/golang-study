package main

import (
	"fmt"
	"time"
)

func fn1(ch chan string) {
	time.Sleep(time.Second * 3)
	ch <- "fn1111"
}

func fn2(ch chan string) {
	time.Sleep(time.Second * 6)
	ch <- "fn22222"
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go fn1(ch1)
	go fn2(ch2)
	//r1 := <- ch1
	//r2 := <- ch2
	//fmt.Println("ch1", r1)
	//fmt.Println("ch2", r2)

	select {
	case r1 := <-ch1:
		fmt.Println("r1:", r1)
	case r2 := <-ch2:
		fmt.Println("r2:", r2)
	}

}
