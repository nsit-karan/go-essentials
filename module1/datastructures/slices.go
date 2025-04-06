package datastructures

import (
	"fmt"
	"reflect"
)

func ProcessSlices() {

	fmt.Println("-------Slice operations -------------- ")

	var slices []int

	slices = append(slices, 10)
	slices = append(slices, 20, 30, 40, 50, 60)

	// elelement at 0, 1, 2 position is referred to in slices1
	// in other words, slices1 is a reference to 0th position
	// and allowed to refer to 0th, 1st, 2nd element only
	slices1 := slices[0:3]

	fmt.Println(slices)
	fmt.Println(slices1)

	// try accessing 4th element - see what happens
	// will throw index out of range runtime panic
	// fmt.Println(slices1[4])

	fmt.Println("Slice initialization 2 ")
	// Append elements to a slice
	// this generates a new slice . refer ex below
	//
	// In the below, arr will hold zero elements
	// while arr1 will now hold 2 elements
	var slice1 []int = []int{}
	slice2 := append(slice1, 100, 200)
	fmt.Println(slice1)
	fmt.Println(slice2)

	// type for both slice1 and slice2 is the same
	// ,i.e , []int
	fmt.Println("Type info about slices")
	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(slice2))

}
