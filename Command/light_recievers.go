package main

import "fmt"

var _ Command = LightOffCommand{}
var _ Command = LightOnCommand{}

type Light struct{}

func (l Light) on() {
	fmt.Println("Light is ON")
}

func (l Light) off() {
	fmt.Println("Light is OFF")
}

type LightOnCommand struct {
	Command
	Light
}

func (LOC LightOnCommand) exec() {
	LOC.Light.on()
}

type LightOffCommand struct {
	Command
	Light
}

func (LOC LightOffCommand) exec() {
	LOC.Light.off()
}
