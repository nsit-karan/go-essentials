package main

import "fmt"

func Add(x int, y int) int {
	fmt.Printf("Sum of the 2 nos is %d\n", x+y)
	return x + y
}

func AddMultiReturn(x int, y int) (bool, int) {
	var sumInt int = x + y

	if sumInt > 10 {
		return true, sumInt
	} else {
		return false, sumInt
	}
}
