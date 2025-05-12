package main

import (
	"strings"
)

func cleanInput(text string) []string {
	s := strings.Fields(text)
	for i, word := range s {
		s[i] = strings.ToLower(word)
	}
	return s
}
