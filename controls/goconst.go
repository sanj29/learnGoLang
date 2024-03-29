package main

import "fmt"

var ar float32 = 0

func main() {

	var radius float32
	var l, b float32

	fmt.Println("Enter Radis")
	fmt.Scanln(&radius)

	fmt.Println("Enter l n b")
	fmt.Scanln(&l, &b)

	area_Of_circle(radius)
	area_of_rectangle(l, b)
}

func area_Of_circle(r float32) {

	const PI = 3.14159
	ar = PI * r * r

	fmt.Println("area of circle ::", int(ar))
}

func area_of_rectangle(l float32, b float32) {
	const HEIGHT int = 100
	const WIDTH int = 200
	ar = l * b
	fmt.Printf("value of area : %d", int(ar))
	fmt.Printf("\n")
}
