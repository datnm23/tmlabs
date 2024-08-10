package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/datnm23/tmlabs/lab03/pkg"
)

var questions = []string{
	"Bạn tên là gì?",
	"Bạn sinh ngày nào?",
	"Bạn làm nghề gì ?",
}

func main() {
	var result []string
	var answer string
	for _, question := range questions {
		fmt.Println(question)
		_, err := fmt.Scan(&answer)

		if err != nil {
			log.Fatal(err)
		}

		result = append(result, question, answer, "")
	}

	data := strings.Join(result, "\n")
	err := pkg.WriteStruct("./person.txt", data)
	if err != nil {
		log.Fatal(err)
	}

}
