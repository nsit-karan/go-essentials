package structs

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) updateUser() {
	u.name = "ludo"
	u.age = 500
}

func (u user) logPerson() {
	fmt.Println(u)
}

func ProcessMethods() {
	var u user = user{name: "hermione", age: 21}
	u.updateUser()
	u.logPerson()
}
