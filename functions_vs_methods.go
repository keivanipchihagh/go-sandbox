package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Weight    float64
	Birthday  time.Time
}

// Function
func UpdateAgeWithFunction(person *Person, age int) {
	person.Age = age
}

// Method
func (person *Person) UpdateAgeWithMethod(age int) {
	person.Age = age
}

func main() {
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Weight:    80.5,
		Birthday:  time.Now(),
	}

	// Never do this!
	person.Age = 40
	fmt.Println(person)

	// Better, but still ugly!
	UpdateAgeWithFunction(&person, 50)
	fmt.Println(person)

	// Now we are talking!
	person.UpdateAgeWithMethod(60)
	fmt.Println(person)
}
