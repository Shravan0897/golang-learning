package main

import "fmt"

func main() {
	const firstName string = "Shravan"
	const lastName = "tomar"
	age := 27
	fmt.Println(firstName, lastName, age)

	const (
		company         string = "xyz"
		designation            = "SDET II"
		experience             = 4
		currentlWorking bool   = true
	)
	fmt.Println(company, designation, experience, currentlWorking)
}
