package main

import (
	"os"
	"strings"
	"fmt"
)
var letters = map[string]string{
	"A": "Apple",
	"B": "Ball",
	"C": "Cat",
}

func getInput(arg string) (string, error) {
	if !strings.HasSuffix(arg, ".txt") {
		return arg, nil
	}

	data, err := os.ReadFile(arg)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func countCharacters(word string) map[rune]int {
	word = strings.ToLower(word)

	counts := make(map[rune]int)

	for _, ch := range word {
		counts[ch]++
	}

	return counts
}

func getLetters(input string) ([]string, error) {
	var result []string

	for _, ch := range input {
		char := string(ch)

		value, ok := letters[char]
		if !ok {
			return result, fmt.Errorf("Error: %s not found", char)
		}

		result = append(result, value)
	}

	return result, nil
}