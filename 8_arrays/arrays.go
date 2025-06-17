package main

import "fmt"

// arrays are numbered sequence of specific length
func main() {
	var nums [5]int //integer array
	nums[0] = 1
	nums[4] = 5
	fmt.Println(len(nums))
	fmt.Println(nums)

	var vals [4]string
	fmt.Println(vals)

	//array declaration with elements
	names := []string{"shravan", "tomar"}
	fmt.Println(names)

	//2d arrays
	relation := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(relation)
}
