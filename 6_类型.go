package main

import "fmt"

func main() {
	//1、布尔类型
	var a bool
	a = false
	fmt.Println(a)
	//2、浮点类型
	var b float32
	b = 3.24
	fmt.Println(b)
	//float64比float32存储小数更准确
	var c float64 = 3.14
	fmt.Println(c)

	//3、字符类型
	var d byte
	d = 97
	fmt.Printf("%c,%d\n", d, d)

	//4、字符串类型
	var e string
	e = "hello"
	fmt.Println(e)

	//5、类型转换
	var f byte
	f = 'a'
	var g int
	g = int(f)
	fmt.Println(g)

	//6、类型别名,int64类型名改为bigint
	type bigint int64
	var h bigint
	h = 100
	fmt.Println(h)
}
