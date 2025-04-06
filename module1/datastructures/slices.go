package datastructures

import "fmt"

func ProcessSlices() {
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
	fmt.Println(slices1[4])
}
