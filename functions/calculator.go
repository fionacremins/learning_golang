package main

import (
	"fmt"
)

// Main Function
func main() {
	var result int
	// Addition
	result = add(6, 9)
	fmt.Println(result)

	// Subtraction
	result = subtract(5, 10)
	fmt.Println(result)

}

// Addition Function
func add(a int, b int) int {

	var ans int
	ans = a + b

	return ans
}

// Subtraction Function
func subtract(a, b int) int {
	var ans int
	ans = a - b

	return ans
}
