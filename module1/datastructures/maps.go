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

	// check for existence of a key
	// accessing a map with a key returns 2 values :
	// the actual value and a boolean to specify if the value was
	// present or not
	value3, ok3 := map3[3]
	value4, ok4 := map3[4]

	if ok3 {
		fmt.Printf("key:%d was found in the above map, value : %s\n",
			3, value3)
	} else if ok4 {
		fmt.Printf("key:%d was found in the above map, value : %s\n",
			4, value4)
	} else {
		fmt.Println("Neither 3 nor 4 was found")
	}

	// delete an entry from a map
	delete(map3, 3)
	fmt.Println("Content of map3 post deletion")
	fmt.Println(map3)

}
