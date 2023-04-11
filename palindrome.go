package main

import "fmt"

/*
Consider an array that contains integers such that each integer is 0 <= x <= 9. The problem statement is to figure out whether this array is a palindrome.

Case 1:
Input: [1, 8, 3, 8, 1]
Output: True

Case 2:
Input: [9, 1, 1]
Output: False
*/

type stack struct {
	value int
}

func checkPalindrome(arr []int) bool {
	for i, j := 0, len(arr)-1; i < len(arr)/2; i++ {
		if arr[i] != arr[j] {
			return false
		}
		j--
	}
	return true
}

func checkPalindromeStack(arr []int) bool {

	var stack []int
	for i := 0; i < len(arr)/2; i++ {
		stack = append(stack, arr[i])
	}

	i := len(arr)/2 - 1
	for j := len(arr)/2 + 1; j < len(arr); j++ {
		if stack[i] != arr[j] {
			return false
		}
		i--
	}
	return true

}

func isPalindrome(head *ListNode) bool {

	slow := head
	fast := head
	var stack []int

	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	len := len(stack)
	for slow != nil {
		if stack[len-1] != slow.Val {
			return false
		}
		slow = slow.Next
		len--
	}

	return true

}

// check palindrome string
func isPalindromeString(str string) bool {
	fmt.Println("Str len:", len(str))
	for i, j := 0, len(str)-1; i < len(str)/2; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
