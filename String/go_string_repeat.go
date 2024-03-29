package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "I like go lang to code"
	fmt.Println(strings.Repeat(str, 4))        // repeats number of time the string
	fmt.Println(strings.Index(str, "l"))       // return the 1st instance of the char
	fmt.Println(strings.Contains(str, "go"))   // return true as go is present in the str
	fmt.Println(strings.Contains(str, "love")) // return false as love not present in the str

	var arr []string = strings.Split(str, " ")
	fmt.Println(len(arr))
	fmt.Println(arr)

	for i, v := range arr {
		fmt.Println("Index : ", i, "value : ", v)
	}
}
