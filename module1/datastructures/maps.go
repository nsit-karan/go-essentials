package datastructures

import (
	"fmt"
)

func ProcessMaps() {
	fmt.Println("\n\n-----------Maps --------------------")
	basicMaps()
}

func basicMaps() {
	map1 := make(map[int]string)
	map1[0] = "stone"
	fmt.Println(map1)

	// initialize a map using map literal to empty map
	map2 := map[int]string{}
	fmt.Println(map2)

	var arr []int = []int{}
	arr1 := append(arr, 100)
	fmt.Println(arr)
	fmt.Println(arr1)

}
