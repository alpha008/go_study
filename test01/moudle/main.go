package main

import (
	"moudle/route"
	"fmt"
)

func main() {
	sum := route.Add(100, 300)
	sub := route.Sub(100, 300)

	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)

	sum, avg := route.Calc(100, 300)
	fmt.Println("sum=", sum)
	fmt.Println("avg=", avg)
}
