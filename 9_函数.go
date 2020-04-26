package main

import "fmt"

func main() {
	process()
	function(2)
}

//
func process() {
	a := 1
	fmt.Println(a)
}
func function(a int) {
	fmt.Println(a)
}
