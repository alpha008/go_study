package main

import(
	"fmt"
	"time"
)
func test_goroute(a int) {
	fmt.Println(a)
}

func loop() {

	for i := 0; i < 100; i++ {
		go test_goroute(i)
	}
	time.Sleep(time.Second)
}
func time_test() {
	 fmt.Println("hello go")
	//go fmt.Println("hello 零声学院") // 在函数内使用前缀go不会执行？？？
	fmt.Println("hello 零声学院") // 在函数内使用前缀go不会执行？？？
	//time.Sleep(time.Second)
}
func main() {
	loop()
	//time_test()
	//fmt.Println("hello world")
	//fmt.Println("零声学院")
}