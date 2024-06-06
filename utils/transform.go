package utils

import (
	"strings"
	"unicode"
)

func Snake(v string) string {
	var result []rune
	for i, r := range v {
		if unicode.IsUpper(r) {
			// Add an underscore before the uppercase letter (except for the first character)
			if i != 0 {
				result = append(result, '_')
			}

			// Convert the uppercase letter to lowercase
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func Pascal(v string) string {
	words := strings.Split(v, "_")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	return strings.Join(words, "")
}
