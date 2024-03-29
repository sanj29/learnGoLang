package main

import (
	"fmt"
	"strconv"
)

func main() {

	var i int = 10
	var f float64 = 6.44
	var str1 = "101"
	var str2 = "10.123"

	fmt.Println(float64(i))
	fmt.Println(int(f))

	newint, _ := strconv.ParseInt(str1, 0, 64)

	fmt.Println(newint)
	newfloat, _ := strconv.ParseFloat(str2, 64)

	fmt.Println(newfloat)

}
