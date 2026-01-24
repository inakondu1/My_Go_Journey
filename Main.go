package main

import "fmt"

func main() {
	var fullName string = "Ina Hadiza Isah"
	var age int = 31
	// let callculate how many years till i turn 50.
	YearsToFifty := 50 - age
	fmt.Println("Hello, World!")
	fmt.Println("my name is", fullName)
	fmt.Println("i am ", age, "years old and am learning go!")
	fmt.Println("in the next ", YearsToFifty, " i will turn 50")

	// this will check if i am 30 or younger

	if age >= 30 {
		fmt.Println("i am old enough")
	} else {
		fmt.Println(" i am younger")
	}
}
