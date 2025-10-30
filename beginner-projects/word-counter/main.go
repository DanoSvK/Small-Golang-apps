package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	text, err := loadDataFromFile("text.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	formattedStr := strings.Fields(text)
	wordsCount := len(formattedStr)
	fmt.Println(wordsCount)
}

func loadDataFromFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("could not open file: %w", err)
	}
	if len(content) == 0 {
		return "", fmt.Errorf("The file is empty")
	}

	return string(content), nil
}
