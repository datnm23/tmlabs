package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileContents, err := os.ReadFile("./contents.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fileContents))
}
