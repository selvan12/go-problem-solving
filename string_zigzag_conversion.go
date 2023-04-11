package main

import "fmt"

/*
Zigzag Conversion
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 1:
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"

Example 2:
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I

Example 3:
Input: s = "A", numRows = 1
Output: "A"
*/

func convert(s string, numRows int) string {
	var result string

	zigZag := make([]string, numRows)

	for i, j := 0, 0; i < len(s); i++ {
		zigZag[j] += string(s[i])
		temp1 := i / (numRows - 1)
		temp2 := temp1 % 2
		//if (i / (numxRows - 1))%2 == 0 { // same as above
		if temp2 == 0 {
			j++
		} else {
			j--
		}
	}
	fmt.Println("zigZag: ", zigZag)

	for _, str := range zigZag {
		result += str
	}

	return result
}
