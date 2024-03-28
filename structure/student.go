package main

import (
	"fmt"
)

type Sudent struct {
	rollnum int
	name    string
	marks   int
}

func main() {

	var s1 Sudent = Sudent{101, "Sanjay", 98}
	var s2 Sudent = Sudent{name: "Bhavya", marks: 99, rollnum: 102}
	fmt.Println(s1)
	fmt.Println(s1.name)

	fmt.Println(s2)
	fmt.Println(s2.name)

}
