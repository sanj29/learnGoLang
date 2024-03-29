package main

import (
	"fmt"
)

func main() {

	var input int

Loop:
	fmt.Printf("You are not eligible to vote \n")
	fmt.Printf("Enter your age :")
	fmt.Scanln(&input)

	if input > 17 {

		fmt.Println("You can Vote")
	} else {
		goto Loop
	}

}
