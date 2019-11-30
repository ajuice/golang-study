package main

import (
	"fmt"
	"sync"
)

func mapTest() {
	// map
	// 普通声明
	m1 := map[string]string{}
	m1["name"] = "郭峰"
	m1["age"] = "18"

	fmt.Println("-------", m1["name"], m1["age"])
}

// 使用make创建
type Person struct {
	Name string
	Age  int
}

func mapTest1() {
	var m map[string]Person
	m = make(map[string]Person)
	m["1"] = Person{
		Name: "呵呵",
		Age:  19,
	}
	p, ok := m["1"]
	fmt.Println(p)
	fmt.Println(ok)
}

// map 的读是线程安全的，读写线程不安全
func mapTest2() {
	var sence sync.Map

	sence.Store("name", "张三")
	sence.Store("age", "22")

	fmt.Println(sence.Load("name"))
	fmt.Println(sence.Load("age"))

	sence.Range(func(key, value interface{}) bool { // 匿名函数
		fmt.Println("key", key, "value", value)
		return true
	})
}
func main() {
	//mapTest()
	//mapTest1()
	mapTest2()
}
