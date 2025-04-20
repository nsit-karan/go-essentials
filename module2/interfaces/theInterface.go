package interfaces

import "fmt"

type Car string
type Motorcycle string

type Drive interface {
	Steer()
}

func (c Car) Steer() {
	fmt.Println("Car can steer")
}

func (m Motorcycle) Steer() {
	fmt.Println("Motorcycle can steer")
}

func canSteer(d Drive) {
	d.Steer()
}

func ProcessInterfaces() {
	canSteer(Car("Car"))
	canSteer(Motorcycle("MotorCycle"))
}
