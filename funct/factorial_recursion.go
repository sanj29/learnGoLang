package main

import "fmt"

func main() {

	fmt.Println(cal_factorial(5))
}

func cal_factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * (cal_factorial(num - 1))
	}
}
