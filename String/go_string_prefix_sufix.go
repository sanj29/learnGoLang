package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "I love Go lang"
	fmt.Println(strings.HasPrefix(str, "I"))
	fmt.Println(strings.HasPrefix(str, "i"))
	fmt.Println(strings.HasSuffix(str, "lang"))
	fmt.Println(strings.HasSuffix(str, "Go"))

}
