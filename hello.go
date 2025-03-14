package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("hello world")
	fmt.Println('A')

	// booleans

	// typed
	var boolA bool = true

	// untyped
	boolB := false
	fmt.Printf("first boolean : %t, second boolean : %t\n", boolA, boolB)

	// integers, floats
	var intA int = 10
	var floatA float64 = 99.99
	fmt.Printf("int is %d, float is %f\n", intA, floatA)

	// find type of the variable
	fmt.Println(reflect.TypeOf(3.1416))

	// initialize 2 variables together
	var name1, name2 string
	name1, name2 = "harry", "ron"
	fmt.Printf("%s, %s\n", name1, name2)

}
