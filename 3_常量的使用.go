package main

import "fmt"

func main() {
	//定义常量
	const a int = 20
	fmt.Println(a)
	//自动推导没有用:=
	const b = 10
	fmt.Println(b)

}
