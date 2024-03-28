package main

import (
	"fmt"
	"math"
)

func mathOps(x float64) float64 {

	var res float64
	res = math.Sqrt(x)

	return res
}

func power(x, y float64) float64 {

	var res float64
	res = math.Pow(x, y)
	return res
}

func mathOpsTest() {

	var a float64 = 3
	var b float64 = 5
	fmt.Println("Using math funcstions")

	fmt.Println("Sqrt of 9 is :", mathOps(9))
	fmt.Println("")
	fmt.Printf("%g to power %g is : %g", a, b, power(a, b))
	fmt.Println()

}
