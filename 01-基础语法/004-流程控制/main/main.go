package main

import (
	"errors"
	"fmt"
)

// if判断
func coon() error {
	return errors.New("这是一个错误")
}

func ifCondition() {
	i := 3
	if i == 3 {

	}

	if i := 3; i == 3 { // 初始化与判断写在一起

	}

	if err := coon(); err != nil { // 特殊用法 , err != nil是判断

	}
}

// 分支语句switch
func switchFunc(num int) {
	// 里面默认写了 break
	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("222")
	default:
		fmt.Println("default")
	}

	switch num {
	case 1:
		fmt.Println("1")
	case 2:
		fallthrough // 不跳出
	case 3:
		fmt.Println("3")
	default:
		fmt.Println(4)

	}

}

// 流程控制	for
func forCondition() {
	var i int
	for ; ; i++ {
		if i > 10 {
			fmt.Println("i", i)
			break
		}
		fmt.Println("当前循环", i)
	}

	//for condition {} 类似	while

	//for {} 死循环
}

// 跳出循环
// break 用于函数类跳出当前的 for, switch select 语句的执行
// continue 用于跳出 for循环的本次迭代
// goto 可以推出多层循环

func main() {
	//switchFunc(2)
	forCondition()
}
