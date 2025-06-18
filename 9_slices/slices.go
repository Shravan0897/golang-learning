package main

import "fmt"

func main() {
	// 1st way of creating a slice
	var nums = make([]int, 0, 5) //1st argument is array, 2nd is intial length, 3rd is initial maximum capacity
	nums = append(nums, 1, 2, 3)
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
	// 2nd way of creating a slice
	names := []string{"Tom", "Jerry"}
	names = append(names, "Shravan", "Aditya")
	fmt.Println(names)
	fmt.Println(len(names))
	fmt.Println(cap(names))

}
