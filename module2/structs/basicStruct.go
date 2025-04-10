package structs

import (
	"fmt"
	"reflect"
)

// anonymous struct where name can't be reused , i.e, cannot create
// another instance of name
// an alternate way of defining the below would be:
// name := struct {name string .....}

var name struct {
	name   string
	age    int
	salary float32
}

func AnonymousStructs() {
	name.name = "harry potter"
	name.age = 20
	name.salary = 12.99

	fmt.Println(name)
	fmt.Println(reflect.TypeOf(name))

	// Alternate way of declaring the above
	name2 := struct {
		name   string
		age    int
		salary float32
	}{
		name:   "hermione",
		age:    20,
		salary: 12.99,
	}

	name2.name = "ron"
	fmt.Println(name2)
	fmt.Println(reflect.TypeOf(name2))
}

// Named structs

type person struct {
	name string
	age  int
}

func processNamedStructs() {
	var personInstance person = person{
		name: "hedwig",
		age:  5,
	}

	fmt.Println(personInstance)
}

func ProcessStructs() {

	fmt.Println("-----Anonymous structs--------")
	AnonymousStructs()

	fmt.Println("-----Named structs------------")
	processNamedStructs()
}
