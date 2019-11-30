package main

import "fmt"

// 数组长度固定， 不可改变

func arrayTest() {
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	for i := 0; i < len(arr); i++ {
		fmt.Println("key", i, "value: ", arr[i])
	}

	for key, value := range arr {
		fmt.Println("key:", key, "value:", value)
	}

	// 数组其它的声明
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr1:", arr1)
	arr2 := [...]int{1, 2, 3, 100} // ... 自动扩展长度
	fmt.Println("arr2:", arr2)

	fmt.Println(arr[:])   // 所有元素
	fmt.Println(arr[:2])  // 右不包含
	fmt.Println(arr[2:3]) // 右不包含
	fmt.Println(arr[2:])  // 左不包含

}

func main() {
	arrayTest()
}
