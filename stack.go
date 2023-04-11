package main

/*
Implement a stack in Golang

A stack is an ordered data structure that follows the Last-In-First-Out (LIFO) principle.
*/

type Stack []int

func (stack *Stack) isEmpty() (bool, int) {
	len := len(*stack)
	if len == 0 {
		return true, len
	}
	return false, len
}

func (stack *Stack) push(val int) {
	*stack = append(*stack, val)
}

func (stack *Stack) pop() int {
	ret := 0
	if flag, len := stack.isEmpty(); !flag {
		ret = (*stack)[len-1]
		*stack = (*stack)[:len-1]
	}
	return ret
}
