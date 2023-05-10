package main

import "fmt"

/*
An example of the Builder pattern in Go.

In this example we will have a builder that builds a car.
*/

type Car struct {
	Seats  int
	Engine string
	GPS    bool
	Weels  int
	Color  string
}

type CarBuilder interface {
	Build() *Car
	SetSeats()
	SetEngine()
	SetGPS()
	SetWeels()
	SetColor()
}

type CarDirector struct {
	builder CarBuilder
}

func (cd *CarDirector) SetBuilder(builder CarBuilder) {
	cd.builder = builder
}

func (cd *CarDirector) Build() *Car {
	return cd.builder.Build()
}

func NewCarDirector(builder CarBuilder) *CarDirector {
	return &CarDirector{builder: builder}
}

func main() {
	sportsBuilder := NewSportsCarBuilder()
	rvBuilder := NewRVCarBuilder()

	sportsDirector := NewCarDirector(sportsBuilder)
	rvDirector := NewCarDirector(rvBuilder)

	sportsCar := sportsDirector.Build()
	rvCar := rvDirector.Build()

	println("Sports car:")
	println("Seats:", sportsCar.Seats)
	println("Engine:", sportsCar.Engine)
	println("GPS:", sportsCar.GPS)
	println("Weels:", sportsCar.Weels)
	println("Color:", sportsCar.Color)

	fmt.Println()

	println("RV car:")
	println("Seats:", rvCar.Seats)
	println("Engine:", rvCar.Engine)
	println("GPS:", rvCar.GPS)
	println("Weels:", rvCar.Weels)
	println("Color:", rvCar.Color)
}
