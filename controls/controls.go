package main

import "fmt"

func numCompare() {

	var x int = 7
	var y int = 7
	var z int = 7

	if (x > y) && (x > z) {
		fmt.Printf("%d is greater", x)
	} else if (y > z) && (y > x) {
		fmt.Printf("%d is greater\n", y)

	} else if (z > x) && (z > y) {
		fmt.Printf("%d is greater\n", z)
	} else if x == y || x == z || z == x {
		fmt.Printf("All numbers are equal\n")
	}

}
