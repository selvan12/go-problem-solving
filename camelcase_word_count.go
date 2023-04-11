package main

import (
	"fmt"
	"unicode"
)

/*
CamelCase:
Given string, determine the number of words in string.

Example:
There are  words in the string: 'one', 'Two', 'Three'.
*/
func camelCaseWordCount() {
	s := "sampleCamelCaseString"

	var out []string
	var start, count int
	for index, val := range s {
		if unicode.IsUpper(val) {
			out = append(out, s[start:index])
			count++
			start = index
		}

		// other way of checking rune
		if 'A' <= val && val <= 'Z' {
			fmt.Println("Capital letter found: ", string(val))
		}
	}

	if start < len(s) {
		out = append(out, s[start:])
		count++
	}

	fmt.Printf("Word Count: %d, Output: %s\n", count, out)
}

/*
Output
Capital letter found:  C
Capital letter found:  C
Capital letter found:  S
Word Count: 4, Output: [sample Camel Case String]
*/
