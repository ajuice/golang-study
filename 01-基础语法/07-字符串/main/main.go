package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

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

// 修改字符串，虽然字符串不能被修改， 但是可以使用其它方式修改
func modifyString() {
	// 方式一 利用byte数组
	str := "hello"
	tempStr := []byte(str)
	tempStr[0] = 'c'
	resultStr := string(tempStr)
	fmt.Println("resultStr:", resultStr)

	// 方式二 利用切片
	tempStr2 := str[1:]
	resultStr2 := fmt.Sprint("s", tempStr2)
	fmt.Println("result2", resultStr2)
}

// len
func lenString() {
	str := "hello"
	str1 := "中国，"
	fmt.Println(len(str))
	fmt.Println(len(str1))                    // 9
	fmt.Println(utf8.RuneCountInString(str1)) // 3 golang中字符串都以 UTF-8 保存，每个中文字符串占3个字节
	//RuneCountInString 可获取真实长度
}

// 字符串遍历
func ergodicString() {
	str := "hello"
	// 方法一
	for i := 0; i < len(str); i++ {
		fmt.Printf("value:%d, %c\n", i, str[i])
	}
	// 方法二
	for key, value := range str {
		fmt.Printf("key: %d, value: %c\n", key, value)
	}

	// 注意：由于上述len()函数本身原因，Unicode字符遍历需要使用range。
}

// 类型转换
func typeTransform() {
	num := 12
	fmt.Println("--------")
	fmt.Printf("%T \n", string(num))
}

// 字符串拼接
func joinString() {
	str1 := "hello"
	str2 := "world"

	var stringBufter bytes.Buffer
	stringBufter.WriteString(str1)
	stringBufter.WriteString(str2)
	resStr := stringBufter.String()
	fmt.Printf("%s\n", resStr)
}

func StringMethods() {
	str := "hello world"
	slice := []string{"whelcome", "to", "china"}
	fmt.Println("e 的位置", strings.Index(str, "e")) //查找索引
	fmt.Println("是否包含 e", strings.Contains(str, "e"))
	fmt.Println("连接 e", strings.Join(slice, ","))
	fmt.Println("替换", strings.Replace(str, "hello", "to", 1)) // 以逗号连接成字符串
	fmt.Println("替换", strings.Split(str, " "))                // 按照空格切分成切片
	fmt.Println("去除头部尾部指定字符串", strings.Trim(str, "h"))
	fmt.Println("去除空格，返回切片", strings.Fields(str))
}

// strconv 转换函数 Append, Parse, Format
func stringConv() {
	// append 将现整数转成字符串添加到现有的字节数组中
	str1 := make([]byte, 0, 100)
	str1 = strconv.AppendInt(str1, 1234, 10) // base进制
	fmt.Println(string(str1))
	str1 = strconv.AppendBool(str1, false)
	fmt.Println(string(str1))
	str1 = strconv.AppendQuote(str1, "dsfsd")
	fmt.Println(string(str1))
	str1 = strconv.AppendQuoteRune(str1, '中')
	fmt.Println(string(str1))

	// Format将其它类型转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.23, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1024)
	fmt.Println(a, b, c, d, e)

	// Parse 将字符串转换成其它类型
	// Parse 系列函数把字符串转换为其他类型
	f, _ := strconv.ParseBool("false")
	g, _ := strconv.ParseFloat("123.23", 64)
	h, _ := strconv.ParseInt("1234", 10, 64)
	i, _ := strconv.ParseUint("12345", 10, 64)
	j, _ := strconv.Atoi("1024")
	fmt.Println(f, g, h, j, i, j)
}

func main() {
	//createdChar()
	//createdString()
	//modifyString()
	//lenString()
	//ergodicString()
	//typeTransform()
	//joinString()
	//StringMethods()
	stringConv()
}
