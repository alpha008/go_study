// 3 reflect反射 基础
package main

import (
	"fmt"
	"reflect"
)
// 反射的是根据值，反推值的类型
func main() {
	var s int = 42
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.ValueOf(s))
	// fmt.Println(reflect.TypeOf(reflect.ValueOf(s)))
}
