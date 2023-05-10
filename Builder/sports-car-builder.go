package main

type SportsCarBuilder struct {
	car *Car
}

func NewSportsCarBuilder() CarBuilder {
	return &SportsCarBuilder{
		car: &Car{},
	}
}

func (scb *SportsCarBuilder) Build() *Car {
	scb.SetColor()
	scb.SetEngine()
	scb.SetGPS()
	scb.SetSeats()
	scb.SetWeels()
	return scb.car
}

func (scb *SportsCarBuilder) SetSeats() {
	scb.car.Seats = 2
}

func (scb *SportsCarBuilder) SetEngine() {
	scb.car.Engine = "V8"
}

func (scb *SportsCarBuilder) SetGPS() {
	scb.car.GPS = false
}

func (scb *SportsCarBuilder) SetWeels() {
	scb.car.Weels = 4
}

func (scb *SportsCarBuilder) SetColor() {
	scb.car.Color = "Red"
}
