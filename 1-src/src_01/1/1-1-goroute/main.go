package main

import(
	"time"
	"fmt"
)
func test_goroute(a int) {
	fmt.Println(a)
}

func main() {

	for i := 0; i < 100; i++ {
		go test_goroute(i)
	}

	time.Sleep(time.Second)
}