package main

import (
	"fmt"
	"sync"
	"time"
)

// 	一 并发的解决方案
/**
1 使用 channel 对 goroutine 间数据交换
2 原子访问 sync/automic
3 互斥锁	sync.Mutex
4 等待组 sync.WaitGroup
*/

// 二 锁
// 互斥锁保证同时只有一个 goroutine 能够访问共享资源 sync.Mutex

func mutexTest() {
	var mutex sync.Mutex
	num := 0
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			num += 1
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second * 3)
	fmt.Println("num = ", num)
}

func main() {

	mutexTest()

	time.Sleep(time.Second * 5)

}
