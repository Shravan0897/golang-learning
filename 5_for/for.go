package main

import "fmt"

// for ->only construct for looping
func main() {
	// for using classic way
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	// for using range
	for j := range 10 {
		fmt.Println(j)
	}
}
