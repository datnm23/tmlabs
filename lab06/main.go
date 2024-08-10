package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/datnm23/tmlabs/lab03/pkg"
)

func main() {
	fmt.Println("Input name of file:")
	var data []byte
	var fileName string
	_, err := fmt.Scan(&fileName)
	fileName = "./" + fileName
	if err != nil {
		log.Fatal(err)
	}

	isExist, err := pkg.CheckFileExisted(fileName)
	if err != nil {
		log.Fatal(err)
	}

	if !isExist {
		fmt.Println("This file doesn't exist")
		os.Exit(0)
	}

	data, err = os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Input number line to print:")
	var y uint
	_, err = fmt.Scan(&y)
	if err != nil {
		log.Fatal(err)
	}

	contents := strings.Split(string(data), "\n")
	if int(y) <= len(contents) {
		fmt.Printf("%d%s%s\n", y, ".", contents[y-1])
	} else {
		fmt.Println("Y is bigger than biggest line of the file")
	}
}
