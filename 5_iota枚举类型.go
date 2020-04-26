package main

import "fmt"

func main() {
	//1、iota常量自动生成器，每一行，自动累加1
	//2、iota给常量赋值使用
	const (
		a = iota //0
		b = iota //1
		c = iota //1
	)
	fmt.Printf("a = %d,b=%d,c=%d\n", a, b, c)
	//3、iota遇到const，重置为0
	const d = iota
	fmt.Println(d)

	//4、可以只写一个iota
	const (
		a1 = iota //0
		b1
		c1
	)
	fmt.Printf("a = %d,b=%d,c=%d\n", a1, b1, c1)
	//5、如果是同一行，值都一样
	const (
		a2, b2, c2 = iota, iota, iota
	)
	fmt.Printf("a = %d,b=%d,c=%d\n", a2, b2, c2)

}
