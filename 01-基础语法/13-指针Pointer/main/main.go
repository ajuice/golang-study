package main

import "fmt"

// 指针，是引用类型
// 牛刀小试
func test() {
	v := 19
	p := &v     // 声明指针p执行地址v
	value := *p // 获取指针p的值

	fmt.Println("v的地址", &v)
	fmt.Println("指针p内存中存的地址", p)
	fmt.Println("指针p的地址", &p)
	fmt.Println("指针p地址指向的值", value)
}

// 使用 new 声明指针
func test1() {
	p := new(*int) // 声明一个int类型的指针
	fmt.Println(p, *p, &p)
}
func main() {
	//test()
	test1()
}
