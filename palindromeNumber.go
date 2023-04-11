package main

import "strconv"

/*
Palindrome Number
Companies
Given an integer x, return true if x is a palindrome, and false otherwise.

Example 1:
Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.

Example 2:
Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

Example 3:
Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
*/

func isPalindromeNumber(x int) bool {
	var reversedNumber int
	temp := x
	for temp > 0 {
		reversedNumber = (reversedNumber * 10) + temp%10
		temp /= 10
	}
	return x == reversedNumber
}

func isPalindromeNumber2(x int) bool {
	var isPalindromeString func(string) bool
	isPalindromeString = func(str string) bool {
		for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
			if str[i] != str[j] {
				return false
			}
		}
		return true
	}

	return isPalindromeString(strconv.Itoa(x))
}
