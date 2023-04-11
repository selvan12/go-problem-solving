package main

import "fmt"

/* Print the numbers in following pattern

55555
4444
333
22
1
*/

// using for loop
func printNumbers(num int) {
	for i := num; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print(i)
		}
		fmt.Println("")
	}
}

// using recursive fuctions
func printValue(count, num int) {
	if count == 0 {
		return
	}
	fmt.Print(num)
	printValue(count-1, num)
}

func printNumbersRecursive(num int) {
	printValue(num, num)
	if num > 0 {
		fmt.Println("")
	}
	if num > 0 {
		printNumbersRecursive(num - 1)
	}
}
