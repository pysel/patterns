// Decorator pattern is a structural pattern
// This pattern is useful when you need to change the action of a piece of code without changing
// that piece of code specifically.
// Idea: to wrap that piece of code into additional functionality
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name        string
	Age         uint64
	Nationality string
}

// Factory function for Person class
func createPerson(name string, age uint64, nationality string) Person {
	return Person{
		Name:        name,
		Age:         age,
		Nationality: nationality,
	}
}

// Now you think: hmmm, it would be great to have a function that creates people with random age.
// For that we need to wrap the createPerson function into our decorator function
func DecoratorCreateRandomAgePerson(name, nationality string, personFactory func(string, uint64, string) Person) Person {
	rand.Seed(time.Now().UnixNano()) // to generate new random numbers each time

	randomAge := rand.Intn(100)
	person := personFactory(name, uint64(randomAge), nationality)
	return person
}

func main() {
	var personWithRandAge Person
	for i := 0; i < 5; i++ {
		personWithRandAge = DecoratorCreateRandomAgePerson("Ruslan", "European", createPerson)
		fmt.Println("Person: ", personWithRandAge.Name, personWithRandAge.Age, personWithRandAge.Nationality)
	}
	// as you can see, all created people have different age field
}
