// my first milestone
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// my second milestone

	var fullName string = "Ina Hadiza Isah."
	// my third milestone
	var age int = 31
	// let callculate how many years till i turn 50.
	YearsToFifty := 50 - age

	fmt.Println("my name is", fullName)
	fmt.Println("I am ", age, "years old and I am learning go!")
	fmt.Println("In the next ", YearsToFifty, "years I will turn 50.")
	// my fouth milestone
	// this will check if i am 30 or younger

	if age >= 30 {
		fmt.Println(" I am old enough to understand Golang.")
	} else {
		fmt.Println(" I am younger")

	}

}

