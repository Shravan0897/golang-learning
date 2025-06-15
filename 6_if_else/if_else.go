package main

import "fmt"

func main() {
	var age int = 18
	if age >= 18 {
		println("This person is an Adult")
	} else {
		fmt.Println("This person is not an Adult")
	}

	// Write a short program that prints each number from 1 to 100 on a new line.
	// For each multiple of 3, print "Fizz" instead of the number.
	// For each multiple of 5, print "Buzz" instead of the number.
	// For numbers which are multiples of both 3 and 5, print "FizzBuzz" instead of the number.

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

}
