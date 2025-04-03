package main

import "fmt"

func TestPointers() {
	PointerBasics()

	var a int = 10
	var b int = 20
	var result int = 0
	PointerAdd(&a, &b, &result)

	fmt.Printf("sum is %d\n", result)

}

func PointerBasics() {
	var value int = 10

	// pointer to an integer
	var ptr *int = &value
	fmt.Printf("Pointer address is %p, value is %d\n", ptr, *ptr)

}

// result gets updated in c
// only for illustration - ideally result should be returned
func PointerAdd(a *int, b *int, c *int) {
	*c = *a + *b
}
