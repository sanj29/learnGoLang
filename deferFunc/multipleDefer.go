package main

import "fmt"

func main() {

	withDefer()
	withoutDefer()

}

func withDefer() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Byee....")
}

func withoutDefer() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("Byee....")
}
