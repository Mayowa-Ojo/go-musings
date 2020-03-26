package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func main() {
	wc.Test(WordCount)
}

// WordCount - checks for duplicate words in a string
func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	wordSlice := strings.Fields(s)

	for _, word := range wordSlice {
		// check if word is a key in map
		count, ok := wordMap[word]

		if !ok {
			wordMap[word] = 1
		}

		wordMap[word] = count + 1
	}

	return wordMap
}
