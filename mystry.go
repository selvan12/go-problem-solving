package main

import "fmt"

// what is the output of mystry(6)
func mystry(n int) {
	if n == 0 || n == 1 {
		return // 				-
	}
	mystry(n - 2)
	fmt.Printf("%d ", n)
	mystry((n - 1))
}

/* Thought process to identify the output:

mystry(6)
	mystry(4)
		mystry(2)
			mystry(0)
		2
		mystry(1)
	4
	mystry(3)
		mystry(1)
	3
	mystry(2)
		mystry(0)
	2
	mystry(1)
6
mystry(5)
	mystry(3)
		mystry(1)
	3
	mystry(2)
		mystry(0)
	2
	mystry(1)
5
mystry(4)
	mystry(2)
		mystry(0)
	2
	mystry(1)
4
mystry(3)
	mystry(1)
3
mystry(2)
	mystry(0)
2
mystry(1)

*/

// Output:
// 2 4 3 2 6 3 2 5 2 4 3 2
