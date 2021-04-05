package main
/*

*/
import "fmt"
//只要实现了speak方法的变量就是speak类型
type speak interface {
	speak()
}
type cat struct {}
type dog struct {}
type person1 struct {}
func (C cat) speak(){
	fmt.Println("cat 抓人")
}
func (D dog) speak(){
	fmt.Println("狗咬人")
}
func (P person1)speak(){
	fmt.Println("打狗，打猫")
}
func Action(action speak){
	action.speak()
}
func main()  {
	var p1 person1
	var d1 dog
	var c1 cat
	Action(d1)
	Action(c1)
	Action(p1)
}

