package main

import "fmt"

func main() {
	//多重赋值
	a, b := 10, 20
	fmt.Printf("a=%d,b=%d\n", a, b)
	//交换变量的值
	a, b = b, a
	fmt.Printf("a=%d,b=%d\n", a, b)
	//匿名变量，丢弃不需要的值，_
	tmp, _ := a, b
	fmt.Println(tmp)

}
