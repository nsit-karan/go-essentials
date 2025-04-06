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

	// initialize a map using map literal with key:value pairs
	map3 := map[int]string{
		1: "stone",
		2: "gold",
		3: "silver",
	}
	fmt.Println(map3)

}
