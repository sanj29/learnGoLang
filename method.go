package main

import "fmt"

func add(x int, y int) int { // last int is return type for this method
	res := x + y
	return res
}

func calc(x, y int) (res, mul int) { // when returing multiple values should be brackets (int, int)
	res = x + y
	mul = x * y

	return
}

/*
func calc(x, y int) (int  int) { // when returing multiple values should be brackets (int, int)

		res = x + y
		mul = x * y

	   return res, mul
	}
*/
func main() {

	sum := add(4, 5)
	sum1, mul := calc(5, 6)

	fmt.Println(" Sum is from calc method", sum1, mul)
	fmt.Println(" Mul is from calc method", mul)

	fmt.Println("sum is ", sum)

}
