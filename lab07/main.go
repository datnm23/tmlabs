package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Input name of file:")
	var data []byte
	var folderName string

	_, err := fmt.Scan(&folderName)
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(folderName)

	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		fmt.Println("Dont't have any file in this folder")
		os.Exit(0)
	}

	fmt.Println("Input key word to search:")
	var keyWord string
	_, err = fmt.Scan(&keyWord)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filepath := filepath.Join(folderName, file.Name())
		data, err = os.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(strings.ToLower(string(data)), strings.ToLower(keyWord)) {
			fmt.Println(file.Name())
			contents := strings.Split(string(data), "\n")
			for i, content := range contents {
				if strings.Contains(content, keyWord) {
					fmt.Printf("%d : %s \n", i+1, content)
				}
			}
		}
	}
}
