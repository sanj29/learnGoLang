package main

import "fmt"

func main() {

	var num [10]int

	for i := 0; i < 10; i++ {

		num[i] = i * 2
	}

	for j := 0; j < 10; j++ {

		fmt.Printf("Element[%d] = %d \n", j, num[j])
	}
}
