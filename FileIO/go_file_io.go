package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Check if the file exists and create a new file if it doesn't
	if _, err := os.Stat("file.txt"); os.IsNotExist(err) {
		file, err := os.Create("file.txt")
		if err != nil {
			log.Fatal(err)
		}
		file.WriteString("Hi... there1")
		file.WriteString("Hi... there2")
		file.WriteString("Hi... there3")
		file.Close()
	} else if err != nil {
		log.Fatal(err)
	}

	// Read the file regardless of the success or failure of the creation
	stream, err := os.ReadFile("file.txt")
	if err != nil {
		log.Println(err)
	}
	readString := string(stream)
	fmt.Println(readString)
}
