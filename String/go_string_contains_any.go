package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Hello", "A"))
	fmt.Println(strings.ContainsAny("Hello", "o & e"))
	fmt.Println(strings.ContainsAny("Hello", ""))
	fmt.Println(strings.ContainsAny("", ""))

	var str string = "\t  \n I love my country  \n\t \r \n"

	fmt.Println("before trim : ", str)
	fmt.Println(strings.TrimSpace(str))
}
