package route
//返回两个参数
func Calc(a int, b int)(int,int) {
    sum := a + b
    avg := (a+b)/2
    return sum, avg
}
