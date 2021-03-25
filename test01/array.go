package main
import "fmt"
func test1() {
	var a [10]int
	//j := 10
	a[0] = 100
	//a[j] = 200
	fmt.Println(a)
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	for index, val := range a { //使用该方法可以便利下标和数组元素
		fmt.Printf("a[%d]=%d\n", index, val)
	}
}

func test3(arr *[5]int) {
	(*arr)[0] = 1000
}

func test2() {
	var a [10]int
	b := a // 值传递
	b[0] = 100
	fmt.Println(a)
}

func main() {

	test1()
	//test2()
	//var a [5]int
	//test3(&a)
	//fmt.Println(a)
}
