package main

import "strings"

func reverseWordInString(s string) string {
	result := ""

	splits := strings.Split(s, " ")
	for i := len(splits) - 1; i > 0; i-- {
		if i != len(splits)-1 {
			result += " "
		}
		result += splits[i]
	}

	return result
}

func reverseWordInString2(s string) string {
	result := ""

	splits := strings.Split(s, " ")
	var revString []string
	for i := len(splits) - 1; i > 0; i-- {
		revString = append(revString, splits[i])
	}

	result = strings.Join(revString, " ")

	return result
}
