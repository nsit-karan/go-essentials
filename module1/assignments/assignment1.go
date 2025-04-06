package assignments

import "fmt"

func CreateMap() map[string]int {
	var mapas map[string]int = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
	}

	return mapas
}

func DeleteMap(mapArg map[string]int) {
	for key, value := range mapArg {
		fmt.Printf("Deleting entry with key : %s, value : %d\n",
			key, value)
		delete(mapArg, key)
	}
}

func PrintMap(mapArg map[string]int) {

	fmt.Println("-----Printing the map contents by hand -------")
	for key, value := range mapArg {
		fmt.Printf("Key: %s, value: %d\n", key, value)
	}
}
