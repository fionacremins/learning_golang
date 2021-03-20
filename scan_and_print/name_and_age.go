package main

import (
	"fmt"
)

func main() {
	fmt.Println("What is your name?: ")

	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is " + name)

	fmt.Println("How old are you?: ")

	var age string
	fmt.Scanln(&age)
	fmt.Println("You are " + age + " years old")

}
