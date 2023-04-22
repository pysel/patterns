package main

import "fmt"

/*
Type: Creational

The Prototype pattern is a creational design pattern that allows creating new instances of an object by cloning existing
instances (prototypes) instead of constructing them from scratch.

This pattern is useful when object creation is expensive or complex and when you need to create many similar objects with different configurations.

The Prototype pattern can help reduce object creation cost, simplify class structure, and enable dynamic object configuration.

In the implementation, a prototype registry will be used to store the prototypes and clone them when needed.
Note, not all implementations of a prototype need to have registry. If the amount of prototypes is small, it is possible to
hardcode them in the client code.
*/

type ColorI interface {
	PrintColor()
	clone() ColorI
}

type Color struct {
	color string
}

func (c *Color) PrintColor() {
	println(c.color)
}

// Imagine cloning is not that expensive as creating a new instance
func (c *Color) clone() ColorI {
	return &Color{c.color}
}

type ColorRegistry map[string]ColorI

func NewRegistry() ColorRegistry {
	c := make(ColorRegistry)

	// imagine that the following operations are very expensive
	c["red"] = &Color{"red"}
	c["blue"] = &Color{"blue"}
	c["green"] = &Color{"green"}

	return c
}

func (c ColorRegistry) GetColor(color string) ColorI {
	if c[color] == nil {
		fmt.Printf("Color %s not registered \n", color)
		return nil
	}
	return c[color].clone()
}

func main() {
	// Set the registry
	registry := NewRegistry()

	// Get the colors from registry using already created prototypes, avoid creating new instances
	red := registry.GetColor("red")
	blue := registry.GetColor("blue")
	green := registry.GetColor("green")
	red2 := registry.GetColor("red")

	// Print the colors
	red.PrintColor()
	blue.PrintColor()
	green.PrintColor()
	red2.PrintColor()
}
