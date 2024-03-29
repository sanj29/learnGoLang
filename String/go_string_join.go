package main

import (
	"fmt"
	"strings"
)

func main() {
	var myds = []string{"a", "b", "c", "d"}
	fmt.Println(strings.Join(myds, "-"))
	fmt.Println(strings.Join(myds, "1122-"))
}
