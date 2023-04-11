package main

import (
	"fmt"
	"strings"
)

/*
Roman to Integer:
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.

Example 1:
Input: s = "III"
Output: 3
Explanation: III = 3.

Example 2:
Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Example 3:
Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	sum := 0
	prev := ""
	input := strings.Split(s, "")
	fmt.Println("String value: ", input)
	for _, val := range input {
		intVal := 0
		fmt.Println("String value: ", val)
		intVal = roman[val]
		switch val {
		case "V", "X":
			if prev == "I" {
				intVal = intVal - 2*roman["I"]
				//prev = ""
			}
		case "L", "C":
			if prev == "X" {
				intVal = intVal - 2*roman["X"]
				//prev = ""
			}
		case "D", "M":
			if prev == "C" {
				intVal = intVal - 2*roman["C"]
				//prev = ""
			}
		}
		prev = val
		sum += intVal
		intVal = 0
	}
	return sum
}
