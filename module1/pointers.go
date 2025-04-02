package main

import "fmt"

func PointerBasics() {
	var value int = 10

	// pointer to an integer
	var ptr *int = &value
	fmt.Printf("Pointer address is %p, value is %d\n", ptr, *ptr)

}
