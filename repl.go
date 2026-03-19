package main

import (
	"fmt"
)

func cleanInput(text string) []string {
	var words []string
	var word string

	for _, c := range text {
		if c == 32 {
			if word != "" {
				words = append(words, word)
				word = ""
			}
			continue
		}
		fmt.Print(text)
		word = word + fmt.Sprintf("%c", c)
		fmt.Println(word)
	}
	if word != "" {
		words = append(words, word)
	}

	fmt.Printf("Resulting slice of strings: %v", words)
	return words
}
