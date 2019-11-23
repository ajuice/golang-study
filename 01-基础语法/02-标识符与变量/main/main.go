package main

import (
	"fmt"
	"os"
)

// go 语言的 25个关键字
/**
if 条件判断 for循环 func 方法或函数 case 条件控制 struct 结构体 import 导入
go 协成 type 类型  chan 管道 defer 延迟 default 默认 package 包定义
map 字典 const 定义常量 else 选择 break 跳出循环 select 协程用 interface 接口
var 变量定义 goto 跳转 range 迭代	return 返回 switch 选择 continue继续 fallthrough穿过
*/

// 保留字
/**
内建常量：
	false true iota nil
内建类型：
	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr
	float32 float64
	complex128 complex64
bool:
	byte rune string error
内建函数：
	make delete complex panic append copy close
	len cap	real imag new recover
*/

// 变量声明
var a int  // 声明一个变量 默认为0
var b = 10 // 声明并初始化 自动推到类型
func variable() {
	// var 声明全局变量
	// := 只能定义在函数内部，
	// 大小写敏感，只能以字母或者下划线开头
	// 推荐驼峰命名

	c := 20 // 初始化，自动推倒类型

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}

// 多变量声明
var e, f string
var a1, a2 = "哼", "哈"
var (
	m int
	n bool
)

func moreVariable() {
	b1, b2 := 1, 2
	fmt.Println("b1, b2", b1, b2)
}

// 变量值交换
func changeVariable() {
	m := "hello"
	n := "world"
	m, n = n, m
	fmt.Println("changeVariable", m, n)
}

// 丢弃值
func abandonVariable() {
	m, n := "hello", "world"
	_, c := m, n
	fmt.Println("abandonVariable:", c)
}

// 声明覆盖
func coverVariable() {
	f, err := os.Open("file")
	out, err := os.Open("file") // 这里的err被上买呢定义的覆盖
	fmt.Println(f, out, err)
}

// 数据的多组书写
const (
	i      = 100
	pi     = 3.14
	prefix = "GO_"
)

var (
	age  int
	name string
	addr string
)

// 关键字 iota 声明初始值为0 每行递增1
const (
	q1 = iota // 0
	q2 = iota // 1
	q3 = iota // 2
	q4 = iota // 3
)
const (
	q5 = iota // 0
	q6        // 1
	q7        // 2
)

const (
	q8           = iota             // 0
	q9, q10, q11 = iota, iota, iota // 同一行，赋值一样 1, 1, 1
	// s = 3 假如这里定义缺省常量，编译报错
)

func main() {
	variable()
	moreVariable()
	changeVariable()
	abandonVariable()
}
