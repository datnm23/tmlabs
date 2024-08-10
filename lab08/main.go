package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// printFormatting prints a single entry with its depth.
func printFormatting(entry string, depth int) {
	indent := strings.Repeat("| ", depth)
	fmt.Printf("%s|--%s\n", indent, entry)
}

// printDirectory prints the directory structure starting from the given path.
func printDirectory(path string) {
	// Internal function to handle recursion with depth.
	var printDirRec func(path string, depth int)

	printDirRec = func(path string, depth int) {
		entries, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		printFormatting(filepath.Base(path), depth)
		for _, entry := range entries {
			fullPath := filepath.Join(path, entry.Name())
			if entry.IsDir() {
				// Recursively print subdirectories
				printDirRec(fullPath, depth+1)
			} else {
				// Print files
				printFormatting(entry.Name(), depth+1)
			}
		}
	}

	// Start the recursive printing with initial depth of 0.
	printDirRec(path, 0)
}

func main() {
	dirPath := os.Args[1]
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run. <%s\n", dirPath)
		return
	}
	if _, err := os.Open(dirPath); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n", dirPath)
		return
	}
	printDirectory(dirPath)
}
