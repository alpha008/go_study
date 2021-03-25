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
	go fmt.Println("hello go")
	go fmt.Println("hello 零声学院") 
	go fmt.Println("hello 零声学院") // 线程还没来得及执行就要退出了
	time.Sleep(time.Second)
}
func main() {
	loop()
	//time_test()
	//fmt.Println("hello world")
	//fmt.Println("零声学院")
}