package main

import "fmt"

// 字符
/**
golang 中没有单独的字符类型(char 类型)，可以用 byte 来创建 字符类型 用单引号包裹的单个字符

字符类型也可以用d%打印为整型
如果我们保存的字符在 ASCII 表的,比如[0-1, a-z,A-Z..]直接可以保存到 byte
如果我们保存的字符对应码值大于 255,这时我们可以考虑使用 int 类型保存
如果我们需要安装字符的方式输出，这时我们需要格式化输出，即 fmt.Printf(“%c”, c1)
字符可以和整型进行运算
*/

func createdChar() {
	c1 := 'c'
	var c2 byte = 'a'
	fmt.Printf("%c, %c", c1, c2)
}

// 字符串，由双引号或者 ``模版字符串包裹起来， 其中 `` 可保留格式, 字符串不可改变

func createdString() {
	str1 := "guofeng"
	str2 := `
	我是大白
`
	fmt.Println(str1, str2)
}

func main() {
	createdChar()
	createdString()
}
