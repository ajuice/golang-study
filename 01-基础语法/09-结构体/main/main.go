package main

import "fmt"

// 结构体的内存是连续的，首字母大写可以是 public , 小写是私有的，针对当前包来说的

type Person struct {
	Name string
	Age  int
}

// 匿名字段
type Student struct {
	Person
	Addr string
}

func main() {
	// 使用结构体
	// 方式一
	var p1 Person
	p1.Name = "郭峰"
	p1.Age = 18
	fmt.Println("--------p1-------", p1.Name, p1.Age)

	// 方式二
	var p2 = new(Person) // new 出来的p2相当于指针， (*p2)
	(*p2).Name = "陶渊明"
	(p2).Age = 19 // 神略了指针操作符
	fmt.Println("--------p2-------", p2.Name, p2.Age)

	//方式三
	p3 := Person{
		Name: "李白",
		Age:  19,
	}
	fmt.Println("--------p3-------", p3.Name, p3.Age)

	p4 := &Person{
		Name: "苏轼",
		Age:  20,
	}
	fmt.Println("--------p4-------", p4.Name, p4.Age)

	// 匿名字段
	var p5 *Student = &Student{}
	p5.Person.Name = "杜甫"
	p5.Person.Age = 19
	p5.Addr = "长安镇"
	fmt.Println("--------p5-------", p5.Name, p5.Age, p5.Addr)
}
