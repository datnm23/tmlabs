package main

import (
	"log"
	"os"

	"github.com/datnm23/tmlabs/lab03/pkg"
)

func main() {
	file := "./content.txt"
	newContent := "Welcome to the world of Go2."
	isExist, err := pkg.CheckFileExisted(file)

	if err != nil {
		log.Fatal(err)
	}

	if !isExist {
		_, err := os.Create(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = pkg.WriteStruct(file, newContent)
	if err != nil {
		log.Fatal(err)
	}

}
