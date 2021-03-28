package main

import (
	"fmt"
	"reflect"
)

type Skills interface {
	reading()
	running()
}

type Student struct {
	Name string "json:name"
	Age  int    "json:age"
}

func (self Student) runing() {
	fmt.Printf("%s is running\n", self.Name)
}
func (self Student) reading() {
	fmt.Printf("%s is reading\n", self.Name)
}
// 设置值的时候，需要这个变量能被修改才可以
func main() {
	stu1 := Student{Name: "darren", Age: 34}
	stu_type := reflect.TypeOf(stu1)
	fmt.Println("stu_type num :",stu_type.NumField())          //2
	fmt.Println(stu_type.Field(0))            //{Name  string  0 [0] false}
	fmt.Println(stu_type.FieldByName("Name")) //{{Age  int  16 [1] false} true
	fmt.Println(stu_type.Field(1))            //{Name  string  0 [0] false}
	fmt.Println(stu_type.FieldByName("Age"))  //{{Age  int  16 [1] false} true
}
