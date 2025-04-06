package datastructures

import (
	"fmt"
	"reflect"
)

func ArraysOps() {
	ArrayInits()
	ArrayInitializations()

	//ArrayIteration()
}

func ArrayInits() {
	var array [3]int
	array[0] = 10

	// initialize the array
	//
	// specifying the type is mandatory for array
	// unlike for primitives
	literal := [2]int{10, 20}
	intBasic := 20

	fmt.Println(literal)
	fmt.Println(intBasic)

}

func ArrayIteration() {
	var array [10]int = [10]int{1, 2, 3, 4, 5}

	fmt.Println("Array iteration - alternative1")
	for i := 0; i < len(array); i++ {
		fmt.Printf("Index:%d, value:%d\n", i, array[i])
	}

	// another alternative
	fmt.Println("\n----Array iteration - alternative2----")
	for i, val := range array {
		fmt.Printf("index%d:value%d \n", i, val)
	}

}

func ArrayInitializations() {

	fmt.Println("\n-----Array Initlization options on the table---------")
	var arr [4]int = [4]int{}
	fmt.Println(arr)

	// you will notice that the type is [4] int
	// ,i.e the size of the array is part of the type
	// in other words arr [4] int will be of a diff type
	// than arr [5] int
	fmt.Println(reflect.TypeOf(arr))
}
