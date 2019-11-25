package main

import (
	"fmt"
	"time"
)

// go routine 关键字 GPM模型与CPS模型

func say(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
	}
}

func main() {
	// 协程先进后出，先做同步，再做异步
	go say("goroutine")
	say("no goroutine")
	time.Sleep(time.Second * 5)
}
