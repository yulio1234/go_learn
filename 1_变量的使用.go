package main

import "fmt"

func main() {
	//1、声明格式
	//2、默认值是0
	//3、{}里变量名是唯一的
	var a int
	fmt.Println("a = ", a)
	//变量可以改变
	a = 10
	fmt.Println("a = ", a)
	//变量初始化并赋值
	var b = 20
	fmt.Println("b = ", b)
	//自动推导类型
	c := 30
	fmt.Printf("c type is %T\n", c)
}
