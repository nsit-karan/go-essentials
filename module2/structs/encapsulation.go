package structs

import (
	"fmt"
	"log"
)

type child struct {
	name  string
	age   int
	adult Adult
}

func ProcessEncapsulation() {
	fmt.Println("---Encapsulation related logic-------")
	var parent Adult = Adult{"bob", 30}
	err := parent.SetAge(40)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(parent)

	// Initializing child (adult present as a field inside child)
	var kid child = child{"spiderman", 15, parent}
	fmt.Println(kid)
}
