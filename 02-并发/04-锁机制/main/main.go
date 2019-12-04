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

type Account struct {
	money int
	lock  *sync.Mutex
}

func (a *Account) Query() {
	fmt.Println("当前账户金额为：", a.money)
}

func (a *Account) Add(num int) {
	a.lock.Lock()
	a.money += num
	a.lock.Unlock()
}

func main() {
	a := &Account{
		0,
		&sync.Mutex{},
	}
	for i := 0; i < 100; i++ {
		go func(num int) {
			a.Add(num)
		}(10)
	}
	time.Sleep(time.Second * 2)
	a.Query()
}
