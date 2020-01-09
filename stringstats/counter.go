package stringstats

import (
	"strings"
)
// CountWords counts the number of words (text separated by a ' ' delimiter) in a string and returns the total number found, represented as a positive integer.
func CountWords(words string) int {
	wordArray := strings.Split(words, " ")
	wordCount := 0
	for i := 0; i < len(wordArray); i++ {
		if wordArray[i] != "" {
			wordCount++
		}
	}
	return wordCount
}