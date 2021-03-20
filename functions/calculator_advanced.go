package main

import (
	"fmt"
	"strconv"
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

	// User Input
	var user string
	fmt.Println("Would you like to add or subtract?: ")
	fmt.Scanln(&user)

	if user == "add" {

		x, y := input()

		ans := add(x, y)

		result := strconv.Itoa(ans)
		fmt.Println("Your answer is " + result)

	} else if user == "subtract" {

		x, y := input()
		ans := subtract(x, y)
		result := strconv.Itoa(ans)
		fmt.Println("Your answer is " + result)
	} else {
		fmt.Println("You did not choose to add or subtract")

	}
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

func input() (int, int) {

	var input1 int
	var input2 int

	fmt.Println("Please enter the first number: ")
	fmt.Scanln(&input1)

	fmt.Println("Please enter the second number: ")
	fmt.Scanln(&input2)

	return input1, input2

}
