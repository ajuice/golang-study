package main

import "fmt"

// 值类型：基本数据类型
// 复合数据类型： 由基本类型通过不同的组合方式构造出的数据类型 数组， slice map 结构体

// 基础类型
/**
整型：int8 uint1
浮点型： float32 float64
复数
布尔型：bool 1字节
数组
结构体 struct
*/

// 引用类型
/**
指针： *
切片：slice
字典：map
函数：func
管道：chan
接口：interface
*/
// golang 没有char类型，可以用byte来保存字符类型

// golang 初始化自带默认值
/**
int		0
int8	0
int32	0
int64	0
uint	0x0
rune	0	//rune的实际类型是 int32
byte	0x0 //byte的实际类型是 uint8
float32 0	// 长度为 4 byte
float64	0	// 长度为 8 byte
bool	false
string	""
*/

// 格式化输出
/**
%%	%字面量
%b	二进制整数值，基数为 2 或者是一个科学技术发表示的指数为2的浮点数
%c	字符型
%d	十进制数值
%e	科学计数法e表示的浮点数或者负数
%E	科学计数法E表述的浮点数或者负数
%f	标准计数法表述的浮点数或者负数
%o	8进制
%p	十六进制表述的一个地址值
%s	字符串
%T	输出值的类型，注意int32和int是两种不通的类型，编译器不会自动转换，需要类型转换
*/

func main() {
	a := 3.0
	fmt.Printf("%T", a) //float64
}
