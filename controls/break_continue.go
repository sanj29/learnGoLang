package main

import "fmt"

func main() {

	var b int = 10
	var a int = 1
	for a < 10 {

		if b%a != 0 {
			fmt.Print("Value of a is ", a, " in continue stmt\n")
			a++
			continue
		}
		fmt.Print("Value of a is ", a, " in break stmt\n")
		a++
		if a > 10 {
			/* terminate the loop using break statement */
			break
		}
	}
}
