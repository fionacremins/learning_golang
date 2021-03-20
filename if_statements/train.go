package main

import (
	"fmt"
	"strings"
)

func main() {

	var train string
	fmt.Println("Where are you going?: ")
	fmt.Scanln(&train)
	train = strings.ToLower(train)

	if train == "cork" {
		fmt.Println("train is going to Cork")
	} else {
		fmt.Println("train is going to Tralee")
	}

}
