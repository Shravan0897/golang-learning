package main

import "fmt"

func main() {
	var firstName string = "Shravan"
	var lastName = "tomar"
	age := 27
	fmt.Println(firstName, lastName, age)

	var (
		company         string = "xyz"
		designation            = "SDET II"
		experience             = 4
		currentlWorking bool   = true
	)
	fmt.Println(company, designation, experience, currentlWorking)
}
