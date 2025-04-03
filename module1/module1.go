package main

import (
	"fmt"
	"reflect"
)

func main() {

	// go type basics
	//goTypes()

	// go conditionals
	//goConditionals()

	// go loops
	//goLoops()

	// functions
	//Add(2, 3)

	isGreater, sum := AddMultiReturn(10, 20)
	fmt.Printf("bool : %t, sum %d\n", isGreater, sum)

	// Pointers
	TestPointers()

}

func goTypes() {
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

func goConditionals() {
	/*
	 * conditionals
	 *
	 * if, else if
	 * &&, ||
	 * ==, !=
	 */
	var str1 string = "harry"
	var str2 string = "harry"

	if str1 == str2 {
		fmt.Println("Strings are equal")
	} else if str1 != str2 {
		fmt.Println("Strings are not really equal")
	} else {
		fmt.Println("Never hit this")
	}

}

func goLoops() {

	fmt.Println("Loop1 - 0 to 9")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("Loop2 - 0, 1, 2")
	var j int = 0
	for j < 10 {
		fmt.Println(j)
		j++

		if j == 3 {
			break
		} else {
			continue
		}
	}
}
