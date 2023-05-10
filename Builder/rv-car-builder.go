package main

type RVCarBuilder struct {
	car *Car
}

func NewRVCarBuilder() CarBuilder {
	return &RVCarBuilder{
		car: &Car{},
	}
}

func (rvcb *RVCarBuilder) Build() *Car {
	rvcb.SetColor()
	rvcb.SetEngine()
	rvcb.SetGPS()
	rvcb.SetSeats()
	rvcb.SetWeels()
	return rvcb.car
}

func (rvcb *RVCarBuilder) SetSeats() {
	rvcb.car.Seats = 6
}

func (rvcb *RVCarBuilder) SetEngine() {
	rvcb.car.Engine = "V6"
}

func (rvcb *RVCarBuilder) SetGPS() {
	rvcb.car.GPS = true
}

func (rvcb *RVCarBuilder) SetWeels() {
	rvcb.car.Weels = 4
}

func (rvcb *RVCarBuilder) SetColor() {
	rvcb.car.Color = "White"
}
