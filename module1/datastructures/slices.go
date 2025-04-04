package datastructures

import "fmt"

func ProcessSlices() {
	var slices []int

	slices = append(slices, 10)
	slices = append(slices, 20, 30, 40)

	fmt.Println(slices)
}
