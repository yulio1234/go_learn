package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum++
	}
	println(sum)
	//通过for打印字符
	str := "abcdefg"
	for i := range str {
		println(i)
	}
	for i, data := range str {

		fmt.Printf("index=%d,data = %c", i, data)
	}
}
