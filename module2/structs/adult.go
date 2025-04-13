package structs

import "errors"

type Adult struct {
	name string
	age  int
}

func (a *Adult) SetAge(newAge int) error {
	if newAge < 1 {
		return errors.New("age cannot be less than 1")
	}

	a.age = newAge
	return nil
}

func (a *Adult) GetAge() int {
	return a.age
}
