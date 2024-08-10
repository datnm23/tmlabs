package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/datnm23/tmlabs/lab03/pkg"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randCharacter() string {
	i := rand.Intn(len(letterRunes))
	return string(letterRunes[i])
}

func main() {
	fmt.Println("Input name of file:")
	var fileName string
	_, err := fmt.Scan(&fileName)

	if err != nil {
		log.Fatal(err)
	}

	fileName = fileName + ".txt"
	fmt.Println("Input X is integer:")

	var x uint
	_, err = fmt.Scan(&x)
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	maxLength := x * 1024 * 1024
	totalLinesFullCharacter := maxLength / 257 // "\n elquals 1 byte so n * 256 + (n-1) = maxLength"
	numCharInLastLine := maxLength % 257
	contents := []string{}

	for lineIndex := uint(0); lineIndex < totalLinesFullCharacter; lineIndex++ {
		var lineContent string
		for i := 0; i < 256; i++ {
			lineContent += randCharacter()
		}
		contents = append(contents, lineContent)
	}

	var lastLineContent string
	for lastLineIndex := uint(0); lastLineIndex < numCharInLastLine; lastLineIndex++ {
		lastLineContent += randCharacter()
	}

	contents = append(contents, lastLineContent)
	result := strings.Join(contents, "\n")
	err = pkg.WriteStruct(fileName, result)
	if err != nil {
		log.Fatal(err)
	}
}
