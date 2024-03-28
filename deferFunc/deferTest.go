package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	a()
}

func a() {
	fmt.Println("i am in method a")
	fmt.Println(" a method starts")
	defer b() // defer will execute the current mehtods then execute next called method.
	fmt.Println(" a method Ends")
	fmt.Printf("current os is : %s \n", runtime.GOOS)
}

func b() {

	var a = 81.0
	fmt.Println("i am in method b")
	fmt.Printf("sqrt of %f is %f \n", a, math.Sqrt(a))

}
