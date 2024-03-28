package main

import "fmt"

var a = 20 // Package level variable

func demo() {
	var a = 10 // funcation level variable

	fmt.Println("In Demo", a)
}

func main() {
	demo()
	fmt.Println("In main", a)

}
