package main

import (
	"fmt"

	interfaces "github.com/module2/interfaces"
	structs "github.com/module2/structs"
)

func main() {
	structs.ProcessStructs()
	structs.ProcessMethods()
	structs.ProcessEncapsulation()

	// Interfaces
	fmt.Println("----Interfaces-------")
	interfaces.ProcessInterfaces()
}
