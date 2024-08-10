package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var vowels = []string{
	"a", "e", "i", "o", "u",
}

var blackListedWords = []string{
	"sex",
	"fuck",
	"drug",
	"kill",
}

var prefixs = []string{
	".",
	",",
}

var suffixs = []string{
	"!",
	".",
	",",
}

func createCensoredWord(blackListedWord string) string {
	for _, vowel := range vowels {
		blackListedWord = strings.ReplaceAll(blackListedWord, vowel, "*")
	}
	return blackListedWord
}

func checkPrefix(word string) bool {
	for _, prefix := range prefixs {
		if strings.HasPrefix(word, prefix) {
			return true
		}
	}

	return false
}

func checkSuffix(word string) bool {
	for _, suffix := range suffixs {
		if strings.HasSuffix(word, suffix) {
			return true
		}
	}

	return false
}

func replaceAllSensitiveWords(contents string) string {
	var newContents []string
	splitRowListOfContents := strings.Split(contents, "\n")

	for _, content := range splitRowListOfContents {
		var newContentInLine []string
		listWordsOfContent := strings.Split(content, " ")
		for _, word := range listWordsOfContent {
			var prefix string
			var suffix string

			if checkPrefix(word) {
				prefix = word[0:1]
				word = word[1:]
			}

			if checkSuffix(word) {
				suffix = word[len(word)-1:]
				word = word[:len(word)-1]
			}
			for _, blackListedWord := range blackListedWords {
				if strings.Compare(strings.ToLower(word), blackListedWord) == 0 {
					word = createCensoredWord(word)
				}
			}

			word = prefix + word + suffix
			newContentInLine = append(newContentInLine, word)
		}

		content := strings.Join(newContentInLine, " ")
		newContents = append(newContents, content)
	}

	return strings.Join(newContents, "\n")
}

func main() {
	contents, err := os.ReadFile("./content.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(replaceAllSensitiveWords(string(contents)))

}
