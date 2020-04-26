package main

import "fmt"

func main() {
	a := 1
	if a == 1 {
		fmt.Println("结果是一")
	} else if a == 2 {
		fmt.Println("结果是二")
	} else {
		fmt.Println("结果什么都不是")
	}
	//条件语句初始化
	if b := 2; b == 2 {
		fmt.Println("结果是二")
	}

	switch a {
	case 1:
		fmt.Println("结果是一")
	case 2:
		fmt.Println("结果是二")
	default:
		fmt.Println("什么都不是")
	}

	switch {
	case false:
		fmt.Println("可以放条件")
	}
}
