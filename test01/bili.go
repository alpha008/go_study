package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)
//结构体
type person struct {
	name string
	age int
	gender string
	hobby[]string
}
func teststruct(){
	var p person
	p.age = 100
	p.name = "alpha"
	p.gender = "男"
	p.hobby = []string{"篮球","足球","乒乓球"}
	fmt.Printf("%T\n",p)
	fmt.Println(p)
}
//pointer
func testpointer(){
	n:=18
	p:=&n
	fmt.Println(p)
	fmt.Printf("%T\n",p)
}
//new
func testnew(){
	var a1 *int
	fmt.Println("a1",a1)
	var a2  = new(int)   //申请一个int类型的空间
	fmt.Println("a2:",a2)
	fmt.Println("*a2",*a2)
	*a2 = 100
	fmt.Println("*a2",*a2)
}
//map
func testmap(){
	//初始化
	var m1 map[string]int
	fmt.Println(m1 == nil) //只是声明，但是没有初始化
	m1 = make(map[string]int,10) // 申请map结构，并初始化10个位置
	m1["alpha"] =100
	m1["spider"] = 200
	fmt.Println(m1)
	//查找
	value,ok := m1["alpha"]  // 返回key对应的value，并且返回是否包含
	if !ok{
		fmt.Println("没有此key")
	} else {
		fmt.Println("alpha",value)
	}
	//遍历key
	for k:= range m1{
		fmt.Print(k)
	}
	//map的遍历 k-v
	for k,v := range  m1{
		fmt.Println(k,v)
	}
}
func testslicemap(){
	rand.Seed(time.Now().Unix())//初始化随机数种子
	var scoreMap = make(map[string]int,200) // 申请大小为200的map
	for i:=0;i<100;i++{
		key:=fmt.Sprintf("stu%2d",i)
		value:=rand.Intn(100)
		scoreMap[key] = value //填充key-value
	}
	// 取出map中所有的元素放入切片keys
	var keys = make([]string,0,200)  // 申请string类型的切片
	for key:=range scoreMap  {      //将map中所有的key放入keys
		keys = append(keys,key)
	}
	// 对切片进行排序
	sort.Strings(keys)  //根据key来进行排序
	for _,key:=range keys{  //根据遍历出来的keys来遍历map
		fmt.Println(key,scoreMap[key])
	}
}
func testslice(){
	var s1[]int  //定义一个存放int的切片
	var s2[]string  // 定义一个存放string类型的切片
	fmt.Println(s1,s2)
	fmt.Println(s1==nil)
	fmt.Println(s2==nil)
	//初始化
	s1 = []int{1,2,3}
	s2 = []string{"shenzhen","shanghai","beijing"}
	fmt.Println(s1,s2)
	fmt.Println(s1==nil)
	fmt.Println(s2==nil)
	fmt.Printf("len(s1):%d ,cap(s1):%d\n",len(s1),cap(s1))
	fmt.Printf("len(s2):%d ,cap(s2):%d\n",len(s2),cap(s2))
	//由数组得到切片
	a1:=[...]int{1,2,3,4,5,6,6}
	s3:=a1[0:4]
	fmt.Println(s3)
	s4:=a1[0:5]
	fmt.Println(s4)
	//make创造切片
	s5:=make([]int,5,10)
	fmt.Printf("len(s5):%d ,cap(s5):%d\n",len(s5),cap(s5))
	s6:=make([]int,1,10)
	fmt.Printf("len(s6):%d ,cap(s6):%d\n",len(s6),cap(s6))
	//切片赋值
	s7:=[]int{1,3,5}
	s8:=s7
	fmt.Println(s7,s8)
	s7[0] = 100
	fmt.Println(s7,s8)
	//切片的遍历
	//for循环遍历
	for i:=0 ; i<len(s7);i++{
		fmt.Println(s7[i])
	}
	//range遍历
	for i,v:=range s7{
		fmt.Println(i,v)
	}
}

func main(){
	//teststruct()
	//testpointer()
	//testnew()
	//testmap()
	//testslicemap()
	//testslice()
}


