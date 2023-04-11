package main

import "fmt"

/* Golang Program To Create Pyramid Pattern
Enter the number
5
        *
      * * *
    * * * * *
  * * * * * * *
* * * * * * * * *
*/
func printPyramid() {
	var num int
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		for j := 0; j < num+i; j++ {
			if j >= num-1 {
				fmt.Print("* ")
			} else if j >= num-i-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func printPyramid2() {
	var num int
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		for j := 0; j < num+i; j++ {
			if j < num-i-1 {
				fmt.Print("  ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}

/*
Enter the number
5
* * * * * * * * *
  * * * * * * *
    * * * * *
      * * *
        *
*/
func printPyramidRev() {
	var num int
	fmt.Scanf("%d", &num)
	for i := num; i > 0; i-- {
		for j := 0; j < num+i-1; j++ {
			if j < num-i {
				fmt.Print("  ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}
}
