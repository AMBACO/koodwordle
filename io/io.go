package io

import (
	"bufio"
	"fmt"
	"os"
)

// LoadWords reads words from a file and returns them as a slice
func LoadWords(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening words file:", err)
		return nil
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 5 { // only 5-letter words
			words = append(words, word)
		}
	}
	return words
}
