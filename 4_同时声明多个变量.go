package main

import "fmt"

func main() {
	var (
		a int
		b float32
	)
	a, b = 10, 3.14
	fmt.Printf("a=%d,b=%f\n", a, b)

	const (
		i = 10
		j = 3.14
	)
	fmt.Println(i, j)
}
