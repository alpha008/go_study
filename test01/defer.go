package main

import "fmt"

var (
	result = func(a1 int, b1 int) int {
		return a1 + b1
	}
)

func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return a1 + b1
	}

	return result(a, b)
}
// 函数存在多个返回值时，使用_取返回值，占位使用，可以不使用
// 定义越靠前 打印越在最后 defer
func main() {
	fmt.Println(result(100, 200))

	var i int = 0
	defer fmt.Println(i)
	defer fmt.Println("second")

	i = 10
	fmt.Println(i)
}
